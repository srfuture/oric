<div align="center">

# Oric マルチプロトコルダウンロードツール

| [中文](../CN/README.md) | [English](../EN/README.md) | [日本語](../JA/README.md) | [한국어](../KO/README.md) |

</div>

> [!NOTE]
> SRFutureによって開発および維持されているOricマルチプロトコルダウンロードツールは、常に新しいプロトコルで更新されています。

> [!CAUTION]
> 現在、amd64アーキテクチャのCPUのみサポートされています。他のアーキテクチャについては、更新を待つか、`oric/cmd/torrent_downloader.cpp` を再コンパイルし、新しい `.so` ライブラリを `oric/lib/` に移動してから、`install.bat` または `install.sh` を再実行してください。

> [!TIP]
> OricはWebクロール機能も提供しており、現在はBilibiliからのビデオダウンロードをサポートしています。

| ダウンロードプロトコル | サポート |
|---------------------|---------|
| HTTP                | ✅      |
| HTTPS               | ✅      |
| FTP                 | ❌      |
| SFTP                | ❌      |
| BitTorrent          | ✅      |
| FTPS                | ❌      |
| WebDAV              | ❌      |
| SCP                 | ❌      |
| Magnet              | ❌      |
| Metalink            | ❌      |

## インストール

- **Windows:**
    1. Gitから:
        - ```cmd
            git clone https://github.com/srfuture/oric
            cd oric
            sudo ./install.sh
    2. ウェブサイトから:
        - 公式ウェブサイトはまだ利用できません

- **Linux:**
    1. Gitから:
        - ```bash
            git clone https://github.com/srfuture/oric
            cd oric
            .\install.bat
    2. ウェブサイトから:
        - 公式ウェブサイトはまだ利用できません

## 使用方法

- URLを使ってダウンロード

    - ```bash
        oric [url_address] -o [download_path]
- BitTorrentを使ってダウンロード

    - ```bash
        oric bt [torrent_file_path] -o [download_path]
- その他のダウンロード
    - bilibiliビデオダウンロード
        - 使用法
        - ``` bash
            oric get [video_url_address] [download_path]
## アップデート

詳細については、[アップデート文書](./update.log.md)を参照してください。
