# Oric 다중 프로토콜 다운로드 도구

<div align="center">

| [中文](../CN/README.md) | [English](../EN/README.md) | [日本語](../JA/README.md) | [한국어](../KO/README.md) |

</div>

> [!NOTE]
> SRFuture에서 개발 및 유지 관리하는 다중 프로토콜 다운로드 도구입니다. 계속해서 프로토콜을 업데이트하고 있습니다.

> [!CAUTION]
> 현재 amd64 아키텍처 CPU만 지원됩니다. 다른 아키텍처가 필요할 경우 업데이트를 기다리거나 직접 `oric/cmd/torrent_downloader.cpp`를 다시 빌드한 후, 새로 생성된 소 라이브러리를 `oric/lib/`에 이동시키고 `install.bat` 또는 `install.sh`를 다시 실행하십시오.

| 다운로드 프로토콜 | 지원 여부 |
|------------------|-----------|
| HTTP             | ✅        |
| HTTPS            | ✅        |
| FTP              | ❌        |
| SFTP             | ❌        |
| BitTorrent       | ✅        |
| FTPS             | ❌        |
| WebDAV           | ❌        |
| SCP              | ❌        |
| Magnet           | ❌        |
| Metalink         | ❌        |

## 설치

- Windows: 
    1. git을 이용한 설치
        - ```cmd
            git clone https://github.com/srfuture/oric
            cd oric
            sudo ./install.sh
    2. 웹사이트를 이용한 설치
        - 현재 공식 웹사이트는 없습니다.

- Linux:
    1. git을 이용한 설치
        - ```bash
            git clone https://github.com/srfuture/oric
            cd oric
            .\install.bat
    2. 웹사이트를 이용한 설치
        - 현재 공식 웹사이트는 없습니다.

## 사용법

- URL로 다운로드

    - ```bash
        oric [URL 주소] -o [다운로드 경로]
- BitTorrent 다운로드

    - ```bash
        oric bt [torrent 파일 경로] -o [다운로드 경로]
- 기타 다운로드

    - > [!TIP]
        > Oric은 현재 bilibili 비디오를 크롤링할 수 있는 기능도 제공합니다.
        
        - 사용법
        - ```bash
            oric get [비디오 URL 주소] [다운로드 경로]