import os
import json
import requests
import subprocess
from selenium import webdriver
from selenium.webdriver.chrome.service import Service
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from time import sleep
from lxml import etree
import pprint
import platform
import argparse

def download_bilibili_video(url, output_dir):
    system = platform.system()
    suffix = ""
    if system == 'Linux':
        machine = platform.machine()
        if machine.startswith('arm'):
            print('ARM芯片的Linux暂不支持此功能')
            return
        else:
            toolspath = "linux_x86"
    elif system == 'Darwin':
        machine = platform.machine()
        if machine.startswith('arm'):
            print('ARM芯片的Mac暂不支持此功能')
            return
        else:
            toolspath = "mac_x86"
    elif system == 'Windows':
        toolspath = "win"
        suffix = ".exe"
    else:
        print('无法识别系统')
        return

    chromedriver_path = f"{toolspath}/chromedriver{suffix}"
    ffmpeg_path = f"{toolspath}/ffmpeg/ffmpeg{suffix}"

    print(f"ChromeDriver path: {chromedriver_path}")
    print(f"FFmpeg path: {ffmpeg_path}")

    # 配置 Selenium
    chrome_options = Options()
    chrome_options.add_argument("--headless")
    chrome_options.add_argument("--no-sandbox")
    chrome_options.add_argument("--disable-dev-shm-usage")
    chrome_options.add_argument("--disable-gpu")
    chrome_options.add_argument("--window-size=1280x1024")

    service = Service(chromedriver_path)

    try:
        driver = webdriver.Chrome(service=service, options=chrome_options)
    except Exception as e:
        print(f"Error initializing WebDriver: {e}")
        return

    try:
        driver.get(url)
        sleep(2)

        page_text = driver.page_source
        driver.quit()
        print(page_text)  # 打印页面内容以帮助调试

        tree = etree.HTML(page_text)

        try:
            title = tree.xpath('//*[@id="viewbox_report"]/h1/text()')[0]
        except IndexError:
            title = tree.xpath('//title/text()')[0]
            print("标题路径不正确，使用了 title 标签")

        try:
            play_info = tree.xpath('/html/head/script[4]')[0].text
            play_info = play_info[20:]
        except IndexError:
            raise ValueError("播放信息未找到")

        play_info_json = json.loads(play_info)
        pprint.pprint(play_info_json)

        video_url = play_info_json['data']['dash']['video'][0]['baseUrl']
        audio_url = play_info_json['data']['dash']['audio'][0]['baseUrl']

        print(f"Video URL: {video_url}")
        print(f"Audio URL: {audio_url}")

        headers = {
            'User-Agent': "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
            'Referer': 'https://www.bilibili.com/',
            'Cookie': "buvid3=5C5D0069-031F-2213-8E11-3B17C971719F69389infoc; b_nut=1688698369; _uuid=7F76CBFD-ADE2-44103-424C-D73D5E9ACC2869255infoc; header_theme_version=CLOSE; CURRENT_FNVAL=4048; buvid4=780B8373-C6A6-6800-F372-7CF18F799AE570981-023070710-7YWVed7pFp%2FuoShCfdfYnQ%3D%3D; DedeUserID=175444232; DedeUserID__ckMd5=b4a676bf5d8afe1c; rpdid=|(k|)mum~~uJ0J'uY))~|uklm; LIVE_BUVID=AUTO5916888971292528; SESSDATA=6b25c9b2%2C1705192174%2Cba23f%2A71bQR5hFBMOt8AXYHjziKE4HOwWw6Ei8wrCIByshPnLAkTd2jwLJy4WYgVkViOyIUPNssSUQAAIAA; bili_jct=e29211bb7e88730fc2bc6691218d247e; sid=858nix09; FEED_LIVE_VERSION=V8; buvid_fp_plain=undefined; hit-new-style-dyn=1; hit-dyn-v2=1; i-wanna-go-back=-1; b_ut=5; fingerprint=b2371c9349b15d5ad60e75cd01f7dc55; buvid_fp=5b9a1047d9ef9ba48290adcd4ba39e58; share_source_origin=copy_web; bsource=share_source_copylink_web; bili_ticket=eyJhbGciOiJIUzI1NiIsImtpZCI6InMwMyIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTY0NzMzNjcsImlhdCI6MTY5NjIxNDEwNywicGx0IjotMX0.D2ixQib5vaXOyxTBLWhIR8KzpbGQloGjxzXDgnOum3E; bili_ticket_expires=1696473307; CURRENT_QUALITY=80; b_lsid=4F245FCD_18AFACA514A; home_feed_column=5; browser_resolution=1552-827; bp_video_offset_175444232=848638555060174904; PVID=1"
        }

        if not os.path.exists(output_dir):
            os.makedirs(output_dir)

        video_path = os.path.join(output_dir, f'{title}.mp4')
        audio_path = os.path.join(output_dir, f'{title}.mp3')

        try:
            with open(video_path, 'wb') as fp:
                fp.write(requests.get(url=video_url, headers=headers).content)
            print(f"Video downloaded to: {video_path}")

            with open(audio_path, 'wb') as fp:
                fp.write(requests.get(url=audio_url, headers=headers).content)
            print(f"Audio downloaded to: {audio_path}")

        except Exception as e:
            print(f"Error downloading video or audio: {e}")
            return

        output_video_path = os.path.join(output_dir, f'{title}_final.mp4')
        ffmpeg_cmd = f'{ffmpeg_path} -i "{video_path}" -i "{audio_path}" -c copy "{output_video_path}"'

        print(f"FFmpeg command: {ffmpeg_cmd}")

        try:
            subprocess.run(ffmpeg_cmd, shell=True, check=True)
            print(f"Final video created at: {output_video_path}")
        except subprocess.CalledProcessError as e:
            print(f"Error running FFmpeg command: {e}")

        try:
            os.remove(video_path)
            os.remove(audio_path)
        except Exception as e:
            print(f"Error removing files: {e}")

    except Exception as e:
        print(f"Error processing page or extracting information: {e}")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="下载 Bilibili 视频并合成音视频文件")
    parser.add_argument('url', help="Bilibili 视频的 URL")
    parser.add_argument('output_dir', help="输出目录")

    args = parser.parse_args()
    download_bilibili_video(args.url, args.output_dir)
