<div align="center">

# Oric多协议下载工具


| [中文](doc/CN/README.md) | [English](doc/EN/README.md) | [日本語](doc/JA/README.md) | [한국어](doc/KO/README.md) |

</div>

> [!NOTE]
> 由诗软未来工作室开发维护更新的的多协议下载工具，不断更新协议中

> [!CAUTION]
> 当前只支持amd64架构CPU，如需其他架构请等待更新或自行重新编译

| 下载协议   | 是否支持 |
|------------|----------|
| HTTP       | ✅       |
| HTTPS      | ✅       |
| FTP        | ❌       |
| SFTP       | ❌       |
| BitTorrent | ✅       |
| FTPS       | ❌       |
| WebDAV     | ❌       |
| SCP        | ❌       |
| Magnet     | ❌       |
| Metalink   | ❌       |

## 安装

- Windows: 
    1. git获取
        - ```cmd
            git clone https://github.com/srfuture/oric
            cd oric
            sudo ./install.sh
    2. 网站获取
        - 暂无官方网站

- Linux:
    1. git获取
        - ```bash
            git clone https://github.com/srfuture/oric
            cd oric
            .\install.bat
    2. 网站获取
        - 暂无官方网站
## 用法

- 使用url下载

    - ```bash
        oric [url地址] -o [下载路径]
- 使用BitTorrent下载

    - ```bash
        oric bt [torrent文件路径] -o [下载路径]
- 其他下载

    - > [!TIP]
      > Oric还提供爬虫功能，目前支持爬取bilibili的视频
    
        - 用法
        - ``` bash
            oric get [视频url地址] [下载路径]
