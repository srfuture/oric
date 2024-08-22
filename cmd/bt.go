package cmd

/*
#cgo CXXFLAGS: -std=c++14
#cgo LDFLAGS: -ltorrent_downloader

#include <stdint.h>
// 声明C++函数
void download_torrent(const char* torrent_file_path,
                      const char* download_rate_limit_str,
                      const char* upload_rate_limit_str,
                      const char* connections_limit_str,
                      const char* max_out_request_queue_str,
                      const char* active_downloads_str,
                      const char* active_seeds_str,
                      const char* active_limit_str,
                      const char* save_path_opt);
*/
import "C"
import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var btTorrent = &cobra.Command{
	Use:   "bt",
	Short: "下载BtTorrent种子文件",
	Long:  `下载BtTorrent(.torrent)种子文件`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if len(args) == 1 {
				optPath, _ := cmd.Flags().GetString("o")
				if optPath != "" {
					fmt.Println("开始下载")
					C.download_torrent(
						toCStringOrNil(args[0]),
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						nil,
						toCStringOrNil(optPath),
					)
				} else {
					fmt.Println(`
用法:
	oric bt [torrent文件路径] [下载速率限制] [上传速率限制] [连接次数限制] [最大外部请求队列大小] [活动下载数量限制] [活动做种数量限制] [活动种子总数量限制] [选项]

选项: 
    -o      下载位置

					`)
					fmt.Println(note("*除了torrent文件路径和下载路径, 其他非必填"))
				}
			} else {
				optPath, _ := cmd.Flags().GetString("o")
				if optPath != "" {
					C.download_torrent(toCStringOrNil(args[0]), toCStringOrNil(args[1]), toCStringOrNil(args[2]), toCStringOrNil(args[3]), toCStringOrNil(args[4]), toCStringOrNil(args[5]), toCStringOrNil(args[6]), toCStringOrNil(args[7]), toCStringOrNil(optPath))
				} else {
					fmt.Println(`
用法:
	iota bt [torrent文件路径] [下载速率限制] [上传速率限制] [连接次数限制] [最大外部请求队列大小] [活动下载数量限制] [活动做种数量限制] [活动种子总数量限制] [选项]

选项: 
    -o      下载位置

					`)
					fmt.Println(note("*除了torrent文件路径和下载路径, 其他非必填"))
				}
			}
		} else {
			note := color.New(color.BgWhite, color.FgBlack).SprintFunc()
			fmt.Println(`
用法:
	iota bt [torrent文件路径] [下载速率限制] [上传速率限制] [连接次数限制] [最大外部请求队列大小] [活动下载数量限制] [活动做种数量限制] [活动种子总数量限制] [选项]

选项: 
    -o      下载位置

			`)
			fmt.Println(note("*除了torrent文件路径和下载路径, 其他非必填"))
		}

	},
}

func init() {
	// 注册 `-o` 选项
	btTorrent.Flags().StringP("o", "o", "", "下载位置")
	// 添加 btTorrent 命令到根命令
	rootCmd.AddCommand(btTorrent)
}

// helper function to convert Go string to C string
func toCStringOrNil(goStr string) *C.char {
	if goStr == "" {
		return nil
	}
	return C.CString(goStr)
}
