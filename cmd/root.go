/*
Copyright © 2024 SRFuture <srfuture2022@gmail.com>

*/
package cmd

import (
	"os"
    "fmt"
	"github.com/spf13/cobra"
	"github.com/fatih/color"

)



var note = color.New(color.BgWhite, color.FgBlack).SprintFunc()

var rootCmd = &cobra.Command{
	Use:   "oric [下载链接] [分段段数]",
	Long: "\n由诗软未来工作室开发并维护的多协议下载工具",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	    if len(args) < 1 {
	        fmt.Println("请提供下载链接。使用-h或-help获取帮助")
	        return
	    }
	    url := args[0]   // 第一个参数是下载链接
	    segments := "1"  // 默认段数
	    if len(args) > 1 {
	        segments = args[1]  // 如果有第二个参数，设置为分段段数
	    }
	    
	    optPath, _ := cmd.Flags().GetString("o")
	    if optPath != "" {
	        // fmt.Printf("开始下载: %s (分段: %s) 保存到: %s\n", url, segments, optPath)
	        httpdownloader(url, optPath, segments)
	    } else {
	        fmt.Println(`
用法:
	oric [下载链接] [分段段数] [选项]
选项:
	-o		下载位置
				`)
	        fmt.Println(note("*分段段数非必填"))
	    }
	},
}




func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// 注册 -o 标志
	rootCmd.Flags().StringP("o", "o", "", "下载位置")

	// 已存在的 toggle 标志
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



