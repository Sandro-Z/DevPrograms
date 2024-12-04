## 写在前面
+ 本文档针对的人群为不熟悉**如何创建go语言微服务**或者是不知道**如何在代码层面对微服务进行修改**的人群
+ 本文档仅可以当做代码开发时**思路**的参考，在具体操作中遇到问题请自行搜索相关内容或者询问资深人员
+ 本文档若有难以理解或者错误之处，请及时联系1370909470@qq.com
## 环境准备
+ ###  编译器的准备
    > 编译器没有强制要求。只要是使用方便，符合个人习惯即可，本文使用[VScode](https://code.visualstudio.com/)来进行示范，本文写于2023年初，使用版本为Vscode2022 <br>
    > 在扩展界面安装相关扩展 <br>
    > ![VScode扩展1](../picture/%E4%BB%8E%E9%9B%B6%E5%BC%80%E5%A7%8B%E7%9A%84%E5%BE%AE%E6%9C%8D%E5%8A%A1/VScode%E6%89%A9%E5%B1%95-1.png)
    > ![VScode扩展2](../picture/%E4%BB%8E%E9%9B%B6%E5%BC%80%E5%A7%8B%E7%9A%84%E5%BE%AE%E6%9C%8D%E5%8A%A1/VScode%E6%89%A9%E5%B1%95-2.png)
+ ### [go语言环境](http://docscn.studygolang.com/)
    > 可以自行在网上搜索相关博客,最后可以在VScode中正常运行helloworld代码即可，这里推荐
    > [一个蛮不错的博客](https://blog.csdn.net/qq_44830881/article/details/123457805)
+ ### [postman](https://www.postman.com/)
    > 一个免费的API调试软件
    > 使用方法在之后会讲到，刚开始安装到自己的电脑即可
+ ### [mysql](https://www.mysql.com/)
    > 基本的底层数据库工具，使用时要求使用sql语句，可以搭配数据库可视化工具开帮助开发，这里推荐一个[Navicate](https://www.navicat.com.cn/download/navicat-premium)
    [](https://www.cnblogs.com/kkdaj/p/16260681.html)
## 操作步骤相关
+ 基本模板的获取以及修改
    1. 下载git.ana中完善的api模板，这里以[api-micro-mail](http://git.ana/xjtuana/api-micro-mail)为例(只有访问到内网且在git.ana上登录后可以访问到该仓库)
    2. **简单修改模板代码，将api-micro-mail换成自身微服务所需要的名称具体操作如下**
       > 首先在所有文件中搜索api-micro-mail，将所有的api-micro-mail改成自己微服务的名称<br>
    3. 创建相关仓库并上传代码
        + 创建仓库
            > ![仓库创建](../picture/%E4%BB%8E%E9%9B%B6%E5%BC%80%E5%A7%8B%E7%9A%84%E5%BE%AE%E6%9C%8D%E5%8A%A1/%E4%BB%93%E5%BA%93%E5%88%9B%E5%BB%BA%E7%95%8C%E9%9D%A2.png)
            > <br> 拥有者选择xjtuana时，该仓库为共有仓库，选择自己用户名时，该仓库为私人仓库，填写仓库名称以及描述后即可创建相关仓库，加入README.md后一次加入上述模板代码后上传到仓库
        + 上传代码
            > 使用[git工具](https://gitforwindows.org/)进行代码的上传
+ 重要文件的作用说明
    + 项目工程文件目录一览(以mail为例)
        - cmd/server
            - main.go 项目功能的主文件，主要功能是负责服务的启动与关闭（一般不需要修改）
        - config
            - config.go 项目工程的基本设置（一般不需要修改）
        - dao 该文件下的文件用于对数据库的具体操作（最底层数据的操作）
            - dao.go 用于打开数据库，数据库配置在config.example.toml中，该文件一般不需要修改
            - XXX.go 用于对数据库进行适用于该微服务的基本操作（使用[GROM语法](https://gorm.io/docs/)）
            - 文件夹，当服务过于复杂或者有相关要求时，将不同的go文件放到不同的目录以便管理
        - dto 该文件下的文件用于对dao下返回的数据进行包装，使其能满足微服务返回的数据标准
            - XXX.go 对返回的数据进行封装
        - http
            - http.go 用于微服务路由的注册，即访问微服务的地址以及对应的函数的定义
            - XXX.go
        - model 该文件下的文件用于对数据库模型的定义
            - model.go 数据库模型的基本定义
            - XXX.go  适用于微服务的模型定义
        - service 该目录下存放业务逻辑
            - service.go
            - XXX.go
        - util（待补充）
        > 剩余的文件部分是项目调试等设置文件，有需要时自行百度即可
    + 工程文件、变量以及函数命名规范（待补充）
+ 基本步骤展示（从api-micro-mail变成自己项目的流程，可以在开发时保留原文件以供参考，但是最后要求删除原有文件）
    1. 注册对应API
        > 在http/http.go文件中的HandlePublic以及HandlePrivate函数中的内容修改为自己微服务的注册路由以及对应的函数<br> 
        ![注册API](../picture/%E4%BB%8E%E9%9B%B6%E5%BC%80%E5%A7%8B%E7%9A%84%E5%BE%AE%E6%9C%8D%E5%8A%A1/%E6%B3%A8%E5%86%8CAPI.png)
    2. dto中创建微服务所需的go文件（命名要求规范）
        > 可以删除项目文件中的其他go文件，初次接触可以仿照已存在文件中的代码
        > <br> type中定义该文件函数绑定数据时所用到的数据类型
        > <br> func中定义数据绑定的具体操作
        <br> ![dto中文件](../picture/%E4%BB%8E%E9%9B%B6%E5%BC%80%E5%A7%8B%E7%9A%84%E5%BE%AE%E6%9C%8D%E5%8A%A1/dto%E6%96%87%E4%BB%B6.png)
    3. service文件下创建微服务所需go文件（命名要求规范）
        > 可以仿照该文件下mail.go中的函数,该文件下调用dao中函数并且使用dto中函数将dao中返回数据绑定为要求规范的数据
        ![service文件](../picture/%E4%BB%8E%E9%9B%B6%E5%BC%80%E5%A7%8B%E7%9A%84%E5%BE%AE%E6%9C%8D%E5%8A%A1/service.png)
    4. model文件下创建微服务对应的go文件
        > 在里面声明自己的数据库模型（一般情况下，数据库中一张表对应一个model）,可以参考该文件下mail.go
        ![model](../picture/%E4%BB%8E%E9%9B%B6%E5%BC%80%E5%A7%8B%E7%9A%84%E5%BE%AE%E6%9C%8D%E5%8A%A1/model.png)
    5. dao文件下创建微服务所需go文件（命名要求规范）
        > 定义相关函数，调用model中的对应文件,对于数据库具体操作可见于[GROM](https://gorm.io/docs/)
    6. 将config.example.toml文件更名为config.toml,并在dsn中配置微服务对应的数据库
        ``` go
        dsn = "user:password@tcp(mariadb:3306)/database?parseTime=True&loc=Local"
        ```
        > user,password是对应数据库的用户名密码，mariadb为数据库地址（一般数据库默认端口为3306），database为数据库库名，其他不用更改
    7. 利用可视化数据库对微服务对应的数据库进行操作，数据库的建立以及修改可以自行百度搜索，可以参考[navicat使用快速入门](https://blog.csdn.net/qq_45069279/article/details/105919312)
## 检测方法以及后续相关
+ postman的使用方法(以查询为例)
    1. 在编译器中启动服务(用以下命令启动该服务)
        ``` shell
        go run cmd/server/main.go
        ```
    2. postman中输入相关api地址以及对应输入
        ![postman测试](../picture/%E4%BB%8E%E9%9B%B6%E5%BC%80%E5%A7%8B%E7%9A%84%E5%BE%AE%E6%9C%8D%E5%8A%A1/postman.png)
        > 有返回值说明该微服务正常启动且有返回值，详细配置可参考[postman中文文档](https://postman.org.cn/)
+ 代码上传规范
    1. commit
        > commit中简短说明对文件的修改，使用中文
    3. 分支规范
        > 要求创建本地分支，在本地将自己的分支与主分支合并之后（包含修改错误）允许提交


## 日志（待删除）
### 等待完成http与service文件的区别，命名规范，util文件
