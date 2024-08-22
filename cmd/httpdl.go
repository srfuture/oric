/*
Copyright © 2024 SRFuture <srfuture2022@gmail.com>

*/

package cmd

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
	"strconv"
	"net/url"
	"time"
)

const maxFileNameLength = 255
const maxRetries = 3
const maxConcurrentDownloads = 50 // 增加并发连接数

// 提取文件名并处理文件系统限制
func getFilePath(urlStr, opt string) (string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	// 处理查询参数，保持文件名
	fileName := path.Base(parsedURL.Path)
	if fileName == "" {
		fileName = "download"
	}

	// 处理文件名长度限制
	if len(fileName) > maxFileNameLength {
		fileName = fileName[:maxFileNameLength]
	}

	// 添加文件路径前缀
	if opt != "" {
		err := os.MkdirAll(opt, os.ModePerm)
		if err != nil {
			return "", err
		}
		fileName = path.Join(opt, fileName)
	} else {
		fileName = path.Join(".", fileName)
	}

	return fileName, nil
}

// 下载文件的一部分，带重试机制
func downloadPart(url string, start, end int64, wg *sync.WaitGroup, progressCh chan<- int64, fileName string, partNum int) {
	defer wg.Done()

	var err error
	for retries := 0; retries < maxRetries; retries++ {
		err = attemptDownloadPart(url, start, end, progressCh, fileName, partNum)
		if err == nil {
			return
		}
		fmt.Printf("部分下载失败，重试中(%d/%d): %v\n", retries+1, maxRetries, err)
		time.Sleep(time.Second * 2)
	}

	fmt.Printf("部分下载最终失败: %v\n", err)
}

func attemptDownloadPart(url string, start, end int64, progressCh chan<- int64, fileName string, partNum int) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))

	client := &http.Client{
		Timeout: time.Minute * 5, // 增加超时时间
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	partFileName := fmt.Sprintf("%s.oric%d", fileName, partNum)
	file, err := os.OpenFile(partFileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Seek(start, io.SeekStart); err != nil {
		return err
	}

	buf := make([]byte, 1024*16) // 增加缓冲区大小
	var totalWritten int64
	const updateInterval int64 = 1024 * 50 // 50 KB
	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			if _, err := file.Write(buf[:n]); err != nil {
				return err
			}
			totalWritten += int64(n)
			if totalWritten >= updateInterval {
				progressCh <- totalWritten
				totalWritten = 0
			}
		}
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
	}
	if totalWritten > 0 {
		progressCh <- totalWritten
	}
	return nil
}

// 下载文件
func downloadFile(url string, parts int, opt string) error {
	fileName, err := getFilePath(url, opt)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: time.Minute * 5, // 增加超时时间
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	totalSize := resp.ContentLength
	if totalSize <= 0 {
		return fmt.Errorf("无法获取文件总大小")
	}

	partSize := totalSize / int64(parts)
	var wg sync.WaitGroup
	progressCh := make(chan int64)
	startTime := time.Now()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		var totalDownloaded int64
		for {
			select {
			case progress, ok := <-progressCh:
				if !ok {
					return
				}
				totalDownloaded += progress
			case <-ticker.C:
				percent := float64(totalDownloaded) / float64(totalSize) * 100
				elapsed := time.Since(startTime).Seconds()
				speed := float64(totalDownloaded) / 1024 / 1024 / elapsed // MB/s

				barWidth := 50
				pos := int(float64(barWidth) * percent / 100)

				fmt.Printf("\033[1A\033[2K\r")
				fmt.Printf("| 下载进度: %3.0f%% | 已下载: %.2f MB | 速度: %.2f MB/s |\n", percent, float64(totalDownloaded)/1024/1024, speed)

				fmt.Print("\033[48;5;255m")
				fmt.Print(strings.Repeat(" ", pos))
				fmt.Print("\033[48;5;235m")
				fmt.Print(strings.Repeat(" ", barWidth-pos))
				fmt.Print("\033[0m")

				fmt.Print("\033[1B")
			}
		}
	}()

	concurrentLimit := make(chan struct{}, maxConcurrentDownloads) // 限制并发数
	for i := 0; i < parts; i++ {
		start := int64(i) * partSize
		end := start + partSize - 1
		if i == parts-1 {
			end = totalSize - 1
		}
		wg.Add(1)
		concurrentLimit <- struct{}{} // 限制并发数
		go func(i int, start, end int64) {
			defer func() { <-concurrentLimit }()
			downloadPart(url, start, end, &wg, progressCh, fileName, i)
		}(i, start, end)
	}

	wg.Wait()
	close(progressCh)
	fmt.Println("\n下载完成!")

	// 合并所有部分
	err = mergeParts(fileName, parts)
	if err != nil {
		fmt.Printf("合并文件失败: %v\n", err)
		return err
	}

	// 校验文件完整性
	if err := verifyChecksum(fileName, "expected_checksum_here"); err != nil {
		return fmt.Errorf("文件校验失败: %v", err)
	}

	return nil
}

// 合并所有部分文件
func mergeParts(fileName string, parts int) error {
	finalFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer finalFile.Close()

	for i := 0; i < parts; i++ {
		partFileName := fmt.Sprintf("%s.oric%d", fileName, i)
		partFile, err := os.Open(partFileName)
		if err != nil {
			return err
		}
		defer partFile.Close()

		_, err = io.Copy(finalFile, partFile)
		if err != nil {
			return err
		}

		os.Remove(partFileName)
	}
	return nil
}

// 验证文件的校验和
func verifyChecksum(fileName, expectedChecksum string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return err
	}

	calculatedChecksum := fmt.Sprintf("%x", hash.Sum(nil))
	if calculatedChecksum != expectedChecksum {
		return fmt.Errorf("校验和不匹配: %s (expected) vs %s (calculated)", expectedChecksum, calculatedChecksum)
	}

	return nil
}

func httpdownloader(url, opt, parts string) {
	part, err := strconv.Atoi(parts)
	if err != nil {
		fmt.Println("错误:", err)
		return
	}
	err = downloadFile(url, part, opt)
	if err != nil {
		fmt.Printf("下载失败: %v\n", err)
	}
}
