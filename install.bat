@echo off
chcp 65001

setlocal

set "current_path=%cd%"

copy "%current_path%\lib\libtorrent_downloader.so" "C:\Windows\System32\"


winget install python

pip install -r "%current_path%\tools\request.txt"


copy "%current_path%\bin\oric_win.exe" "%current_path%\oric.exe" 

echo 正在删除缓存，询问是否删除请选 “Y”

del "%current_path%\bin"

echo 安装完成，请打开环境变量编辑器，双击编辑Path变量，添加值: "%current_path%"

endlocal
