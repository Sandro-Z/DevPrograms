# ivpn client

ivpn 客户端

### 预编译的二进制文件

访问 [GitHub Actions](https://git.ana/dorbmon/ivpn-gui/actions/workflows/build-branch-master.yaml) 页面，点击最新的自动构建，下载与自己平台相对应的 Artifact 即可使用

Windows 平台需要额外下载 [wintun.dll](https://www.wintun.net/) 并将对应架构的动态链接库解压到二进制相同目录下

### 使用方法

```shell
sudo ivpnc --device <tun> --proxy wss://<token>@<addr>/<path>

%when use socks5
sudo ivpnc --device <tun> --proxy wss://<token>@<addr>/<path> -enable-socks5 --socks5-address :1234 -enable-http --http-port 1080
```

具体路由配置可参考 [tun2socks](https://github.com/xjasonlyu/tun2socks/wiki/Examples) 中网卡 IP 及路由配置部分



