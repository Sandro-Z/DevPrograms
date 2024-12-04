# DevPrograms
这个仓库汇总了git.ana中全部的研发部仓库项目。测试项目，没有任何文件的项目，招新Task（ToDoList）以及镜像的仓库没有包括。
同名仓库已经过筛选，保留了最新，最完整的代码。同时清除了无内容的readme files和无意义的测试文档。

下面是有关这些项目的索引，按照首字母顺序排列：

## 项目索引
| 名称 | 作者 | README | Code | 描述 |
|------|------|--------|------|-------|
|api-common|Uebian|[readme](./api-common/README.md)|[file](./api-common)|作用尚不明确。|
|apidoc|tomdang|null|[file](./apidoc)|原网站api文档|
|api-micro-test|ncoder|[readme](./api-micro-test/README.md)|[file](./api-micro-test)|从零开始新建的简单微服务项目（但是readme上写的是micro mail？）|
|Chatbot|null|null|[file](./chatbot)|众多社员一同训练的Q群智能答疑机器人，但好像已经年久失修了……|
|Deland|Shiwei Wang|null|[file](./deland)|作用尚不明确。|
|DNS-config|abdtyx|null|[file](./dns-config)|内网DNS配置文件|
|ivpn-gui|Dorbmon|[readme](./ivpn-gui/README.md)|[file](./ivpn-gui)|ivpn客户端|
|littlebackend|nighttale|[readme](./littlebackend/README.md)|[file](./littlebackend)|gin框架编写的后端|
|Network-manager|ncoder|[readme](./network_manager/README.md)|[file](./network_manager)|作用尚不明确。|
|pub-config|Uebian|null|[file](./pub-config)|似乎是旧内网的配置，可考虑使用该配置重新搭建wireguard|
|QQguildgo|Shiwei Wang|null|[file](./qqguildgo)|作用尚不明确。|
|roles|Liu Qianyu|null|[file](./roles)|作用尚不明确。猜测可能与数据库中社员的roles有关？|
	
## 新人须知

这些是原docs库中一些介绍性的内容，可以让新社员快速熟悉社团的工作环境。

* [内网访问](./docs/内网/新版内网访问文档.md)

* [社团网络架构](./docs/网络架构/network.md)

* [社员福利](./docs/社员福利.md)

* [研发部成员须知](./docs/研发部成员须知.md)

* [镜像站相关](./docs/镜像站相关.md)

## XJTUANA内网的常用服务
|  地址   | 协议  | 服务描述| 所需身份 |
|  ----  | ----  | ---- | ---- |
| [git.ana](http://git.ana) | HTTP | 社团git服务 | 社员 |
| [drone.ana](http://drone.ana) | http | drone CI | 研发部成员 |
| [pma.ana](http://pma.ana) | HTTP | 研发部phpmyadmin服务 | 研发部成员 |
| mariadb.nic1.ana | mariadb | 测试环境MariaDB数据库 | 研发部成员 |
| mariadb.nic5.ana | mariadb | 生产环境MariaDB数据库 | 研发部部长，研发副社长，社长 |
| mariadb.m1.ana | mariadb | 生产环境MariaDB数据库 | 研发部部长，研发副社长，社长 |
| mongodb.m2.ana | mongodb | 生产环境MongoDB数据库 | 研发部部长，研发副社长，社长 |
| [swagger-ui.ana](http://swagger-ui.ana) | http | Swagger文档查看器 | 社员 |
| ~~pan.ana~~ | http | 百度网盘离线下载服务 | 社员 |
| ~~gpt.ana~~ | http | ChatGPT服务 | 社员 |
| 10.58.17.201:1080 | socks5 | 出国代理服务，出口为HK | 社员 |
| *.dev.ana | 多种协议 | 项目的测试环境，其中*替换为对应项目的官方名称 | 研发部成员 |

## 相关站点

* [网管会主页](https://xjtuana.com/)

* [西交官方镜像站](https://mirrors.xjtu.edu.cn/)

* [测速站](https://speed.xjtu.edu.cn/)

* [招新系统](https://apply.xjtuana.cn)

* [ivpn系统](http://ana.xjtu.edu.cn/ivpn/)