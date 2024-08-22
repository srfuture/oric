<div align="center">

# Oric 다중 프로토콜 다운로드 도구

| [中文](../CN/README.md) | [English](../EN/README.md) | [日本語](../JA/README.md) | [한국어](../KO/README.md) |

</div>

> [!NOTE]
> SRFuture에 의해 개발 및 유지되는 Oric 다중 프로토콜 다운로드 도구는 지속적으로 새로운 프로토콜로 업데이트됩니다.

> [!CAUTION]
> 현재는 amd64 아키텍처의 CPU만 지원합니다. 다른 아키텍처의 경우 업데이트를 기다리거나, `oric/cmd/torrent_downloader.cpp`를 재컴파일한 후 새 `.so` 라이브러리를 `oric/lib/`로 이동한 다음 `install.bat` 또는 `install.sh`를 다시 실행하세요.

> [!TIP]
> Oric은 웹 크롤링 기능도 제공하며, 현재 Bilibili에서 비디오 다운로드를 지원합니다.

| 다운로드 프로토콜 | 지원 여부 |
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

## 설치

- **Windows:**
    1. Git에서:
        - ```cmd
            git clone https://github.com/srfuture/oric
            cd oric
            sudo ./install.sh
    2. 웹사이트에서:
        - 공식 웹사이트는 아직 제공되지 않습니다

- **Linux:**
    1. Git에서:
        - ```bash
            git clone https://github.com/srfuture/oric
            cd oric
            .\install.bat
    2. 웹사이트에서:
        - 공식 웹사이트는 아직 제공되지 않습니다

## 사용법

- URL을 사용하여 다운로드

    - ```bash
        oric [url_address] -o [download_path]
- BitTorrent를 사용하여 다운로드

    - ```bash
        oric bt [torrent_file_path] -o [download_path]
- 기타 다운로드
    - bilibili 비디오 다운로드
        - 사용법
        - ``` bash
            oric get [video_url_address] [download_path]
## 업데이트

자세한 내용은 [업데이트 문서](./update.log.md)를 참조하세요.
