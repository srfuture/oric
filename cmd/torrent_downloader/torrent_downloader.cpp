/*
Copyright © 2024 SRFuture <srfuture2022@gmail.com>

*/
#include <libtorrent/session.hpp>
#include <libtorrent/add_torrent_params.hpp>
#include <libtorrent/torrent_handle.hpp>
#include <libtorrent/magnet_uri.hpp>
#include <iostream>
#include <thread>
#include <string>
#include <iomanip> // 用于格式化输出

extern "C" {
    void download_torrent(const char* torrent_file_path,
                          const char* download_rate_limit_str,
                          const char* upload_rate_limit_str,
                          const char* connections_limit_str,
                          const char* max_out_request_queue_str,
                          const char* active_downloads_str,
                          const char* active_seeds_str,
                          const char* active_limit_str,
                          const char* save_path_opt);
}

void download_torrent(const char* torrent_file_path,
                      const char* download_rate_limit_str = nullptr,
                      const char* upload_rate_limit_str = nullptr,
                      const char* connections_limit_str = nullptr,
                      const char* max_out_request_queue_str = nullptr,
                      const char* active_downloads_str = nullptr,
                      const char* active_seeds_str = nullptr,
                      const char* active_limit_str = nullptr,
                      const char* save_path_opt = nullptr) {
    using namespace libtorrent;

    if (!save_path_opt || std::string(save_path_opt).empty()) {
        throw std::invalid_argument("下载路径不能为空");
    }

    int download_rate_limit = download_rate_limit_str ? std::stoi(download_rate_limit_str) : 0;
    int upload_rate_limit = upload_rate_limit_str ? std::stoi(upload_rate_limit_str) : 0;
    int connections_limit = connections_limit_str ? std::stoi(connections_limit_str) : 50000;
    int max_out_request_queue = max_out_request_queue_str ? std::stoi(max_out_request_queue_str) : 10000;
    int active_downloads = active_downloads_str ? std::stoi(active_downloads_str) : 100;
    int active_seeds = active_seeds_str ? std::stoi(active_seeds_str) : 100;
    int active_limit = active_limit_str ? std::stoi(active_limit_str) : 5000;

    settings_pack settings;
    settings.set_str(settings_pack::user_agent, "libtorrent/1.2.11");
    settings.set_int(settings_pack::download_rate_limit, download_rate_limit);
    settings.set_int(settings_pack::upload_rate_limit, upload_rate_limit);
    settings.set_int(settings_pack::connections_limit, connections_limit);
    settings.set_int(settings_pack::max_out_request_queue, max_out_request_queue);
    settings.set_int(settings_pack::active_downloads, active_downloads);
    settings.set_int(settings_pack::active_seeds, active_seeds);
    settings.set_int(settings_pack::active_limit, active_limit);
    settings.set_bool(settings_pack::strict_super_seeding, false);
    settings.set_bool(settings_pack::enable_lsd, true);
    settings.set_bool(settings_pack::enable_upnp, true);
    settings.set_bool(settings_pack::enable_natpmp, true);

    session ses(settings);
    std::string torrent_file_path_str(torrent_file_path);
    add_torrent_params p;
    p.ti = std::make_shared<torrent_info>(torrent_file_path_str);

    // 使用指定的保存路径
    p.save_path = save_path_opt;

    torrent_handle h = ses.add_torrent(p);

    std::cout << "开始下载: " << torrent_file_path << std::endl;

    bool downloading = true;
    while (downloading) {
        torrent_status s = h.status();

        // 计算进度百分比
        int progress = static_cast<int>(s.progress * 100);
        int bar_width = 50; // 进度条宽度
        int pos = bar_width * progress / 100;

        // 清除当前行并返回行首
        std::cout << "\033[2K\r"; // 清除整行并将光标移动到行首

        // 显示信息
        std::cout << "| 下载进度: " << std::setw(3) << progress << "% | ";

        // 计算已下载大小及其单位
        double total_done = static_cast<double>(s.total_done);
        std::string unit;
        if (total_done < 1024) {
            unit = "B";
            std::cout << "已下载: " << std::fixed << std::setprecision(2) << total_done << " " << unit << " | ";
        } else if (total_done < 1024 * 1024) {
            unit = "KB";
            std::cout << "已下载: " << std::fixed << std::setprecision(2) << total_done / 1024 << " " << unit << " | ";
        } else if (total_done < 1024 * 1024 * 1024) {
            unit = "MB";
            std::cout << "已下载: " << std::fixed << std::setprecision(2) << total_done / (1024 * 1024) << " " << unit << " | ";
        } else if (total_done < 1024 * 1024 * 1024 * 1024) {
            unit = "GB";
            std::cout << "已下载: " << std::fixed << std::setprecision(2) << total_done / (1024 * 1024 * 1024) << " " << unit << " | ";
        } else if (total_done < 1024 * 1024 * 1024 * 1024 * 1024) {
            unit = "TB";
            std::cout << "已下载: " << std::fixed << std::setprecision(2) << total_done / (1024 * 1024 * 1024 * 1024) << " " << unit << " | ";
        } else {
            unit = "PB";
            std::cout << "已下载: " << std::fixed << std::setprecision(2) << total_done / (1024 * 1024 * 1024 * 1024 * 1024) << " " << unit << " | ";
        }

        // 显示下载速度
        std::cout << "下载速度: " << s.download_rate / 1024 << " KB/s |";

        // 输出换行符以将进度条移到文本下方
        std::cout << "\n";

        // 绘制进度条背景
        std::cout << "\033[48;5;255m"; // 设置背景色为白色
        std::cout << std::setw(pos) << " "; // 填充进度条背景

        // 绘制未完成的部分
        std::cout << "\033[48;5;235m"; // 设置背景色为灰色
        std::cout << std::setw(bar_width - pos) << " "; // 填充未完成的部分

        std::cout << "\033[0m"; // 重置颜色

        // 刷新输出流
        std::cout.flush();

        // 清除进度条行并将光标移到行首
        std::cout << "\033[1A\033[2K\r"; // 向上移动一行，清除该行内容，并返回行首

        if (s.progress == 1.0) {
            downloading = false;
        }

        std::this_thread::sleep_for(std::chrono::seconds(1));
    }

    // 打印完成消息
    std::cout << "\n下载完成!" << std::endl;
}
