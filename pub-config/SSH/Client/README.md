# OpenSSH客户端配置

将客户端配置`config`和主机指纹`known_hosts`下载到`~/ssh/`目录中。

```
wget -O ~/.ssh/config http://git.ana/xjtuana/pub-config/raw/branch/master/SSH/Client/config
wget -O ~/.ssh/known_hosts http://git.ana/xjtuana/pub-config/raw/branch/master/SSH/Client/known_hosts
```

然后，修改`config`中的私钥路径为你的私钥文件实际路径。

Windows系统下的路径通常为`C:\Users\{UserName}\.ssh`。

## 连接服务器

```
# 从内网连接
ssh nic2.ana
```

## 重要提示

各服务器的主机指纹已经写在`known_hosts`中，同时服务器指纹校验已经写在配置文件中`StrictHostKeyChecking yes`。

出于安全考虑，请**务必不要**手动关闭指纹校验。
