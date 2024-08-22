#!/bin/bash

current_path=$(pwd)

sudo cp "$current_path/lib/libtorrent_downloader.so" /usr/lib

sudo apt install -y python-is-python3

pip install -r "$current_path/tools/request.txt"

case "$(uname)" in
    Linux)
        mv "$current_path/bin/oric_linux_amd64" "$current_path/oric"

        rm -rf "$current_path/bin"
        ;;
    Darwin)
        mv "$current_path/bin/oric_mac_amd64" "$current_path/oric"

        rm -rf "$current_path/bin"
        ;;
    *)
        echo "不支持此系统"
        exit 1
        ;;
esac

echo "请运行'echo "export PATH=\"$current_path:\$PATH\"" >> ~/.bashrc' 配置环境变量"

echo "安装完成请运行'source ~/.bashrc'或重启以彻底完成安装"
