<div align="center">

# Oric Multi-Protocol Download Tool

| [中文](../CN/README.md) | [English](../EN/README.md) | [日本語](../JA/README.md) | [한국어](../KO/README.md) |

</div>

> [!NOTE]
> Developed and maintained by SRFuture, the Oric multi-protocol download tool is constantly updated with new protocols.

> [!CAUTION]
> Currently only supports amd64 architecture CPUs. For other architectures, please wait for updates or recompile `oric/cmd/torrent_downloader.cpp`, then move the new `.so` library to `oric/lib/`, and finally run `install.bat` or `install.sh` again.

> [!TIP]
> Oric also provides a web crawling feature, currently supporting video downloads from Bilibili.

| Download Protocol | Supported |
|-------------------|-----------|
| HTTP              | ✅        |
| HTTPS             | ✅        |
| FTP               | ❌        |
| SFTP              | ❌        |
| BitTorrent        | ✅        |
| FTPS              | ❌        |
| WebDAV            | ❌        |
| SCP               | ❌        |
| Magnet            | ❌        |
| Metalink          | ❌        |

## Installation

- **Windows:**
    1. From Git:
        - ```cmd
            git clone https://github.com/srfuture/oric
            cd oric
            sudo ./install.bat
    2. From Website:
        - No official website available yet

- **Linux:**
    1. From Git:
        - ```bash
            git clone https://github.com/srfuture/oric
            cd oric
            .\install.sh
    2. From Website:
        - No official website available yet

## Usage

- Download using URL

    - ```bash
        oric [url_address] -o [download_path]
- Download using BitTorrent

    - ```bash
        oric bt [torrent_file_path] -o [download_path]
- Other Downloads
    - bilibili video download
        - Usage
        - ``` bash
            oric get [video_url_address] [download_path]
## Update

Please refer to the [update document](./update.log.md) for details.
