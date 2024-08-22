/*
Copyright © 2024 SRFuture <srfuture2022@gmail.com>

*/
package cmd

import (
    "fmt"
    "os/exec"
	"github.com/spf13/cobra"
)

func run(url, optPath string) {
    // 要执行的命令和参数
    cmd := exec.Command("python","tools/bilibilivideocrawler.py", url, optPath)

    // 捕获命令的标准输出和标准错误
    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    // 打印命令的输出
    fmt.Printf("%s", output)
}


var Crawler = &cobra.Command{
	Use:   "get",
	Short: "爬取视频网站视频, 目前只支持Bilibili",
	Long:  `爬取视频网站视频, 目前只支持Bilibili`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			run(args[0], args[1])
		}else{
			fmt.Printf(`
用法: 
	oric get [视频url] [下载路径]
			
			`)
		}
	},
}

func init() {
	rootCmd.AddCommand(Crawler)
}
