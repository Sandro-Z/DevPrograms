# 内网DNS相关说明

 - 主要DNS服务器：`10.58.11.2` （域名`dns.ana`）
 - 备用DNS服务器：`10.58.12.2` （仅服务器可用）

说明：
 - 负责解析内网域名`*.ana`
 - 内网DNS提供纯净DNS服务，可以正常解析墙外域名
 - 仅支持标准DNS协议(UDP 53)，不支持DNSSEC、DoT和DoH等协议
 - 内网使用mosdns软件提供DNS服务，配置文件地址见`http://git.ana/xjtuana/dns-config`

# 主要服务域名(Service)

 - [开发测试环境管理平台](http://portainer.nic3.ana)

 用途：主要用于开发测试环境的开启、关闭、清理等管理工作

 域名：`http://portainer.nic3.ana`

 - VSCode远程开发环境

 用途：用于个人远程开发调试程序

 域名：`*.vscr.ana`

 - [Git仓库](http://git.ana)

 用途： 存放代码、文档、公共资料、笔记等等
 
 域名：`http://git.ana`

 - [phpMyAdmin](http://pma.ana)
 
 用途：MySQL/MariaDB数据库的图形化管理
 
 域名：`http://pma.ana`
