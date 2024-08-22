# Oric多プロトコルダウンロードツール
<div align="center">

| [中文](../CN/README.md) | [English](../EN/README.md) | [日本語](../JA/README.md) | [한국어](../KO/README.md) |

</div>

> [!NOTE]
> SRFuture（詩軟未来工作室）が開発・保守・更新している多プロトコルダウンロードツールです。プロトコルは継続的に更新されています。

> [!CAUTION]
> 現在、amd64アーキテクチャのCPUのみサポートされています。他のアーキテクチャが必要な場合は、更新を待つか、`oric/cmd/torrent_downloader.cpp` を再パッケージして、新しい.soライブラリを `oric/lib/` に移動し、最後に `install.bat` または `install.sh` を再実行してください。

| ダウンロードプロトコル | サポート状況 |
|----------------------|--------------|
| HTTP                 | ✅           |
| HTTPS                | ✅           |
| FTP                  | ❌           |
| SFTP                 | ❌           |
| BitTorrent           | ✅           |
| FTPS                 | ❌           |
| WebDAV               | ❌           |
| SCP                  | ❌           |
| Magnet               | ❌           |
| Metalink             | ❌           |

## インストール

- Windows:
    1. gitから取得
        - ```cmd
            git clone https://github.com/srfuture/oric
            cd oric
            sudo ./install.sh
    2. ウェブサイトから取得
        - 現在、公式ウェブサイトはありません

- Linux:
    1. gitから取得
        - ```bash
            git clone https://github.com/srfuture/oric
            cd oric
            .\install.bat
    2. ウェブサイトから取得
        - 現在、公式ウェブサイトはありません

## 使い方

- URLでダウンロード

    - ```bash
        oric [urlアドレス] -o [ダウンロードパス]
- BitTorrentでダウンロード

    - ```bash
        oric bt [torrentファイルパス] -o [ダウンロードパス]
- その他のダウンロード

    - > [!TIP]
        > Oricはビリビリ動画のクローリング機能も提供しています。
    
        - 使い方
        - ``` bash
            oric get [動画urlアドレス] [ダウンロードパス]
