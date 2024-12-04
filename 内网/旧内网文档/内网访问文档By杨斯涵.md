# 猪都能明白的如何访问XJTUANA内网
1. 群文件下载 network.7z
2. 前往
    - 校外访问 [社团首页 - 学生网络管理协会 - 西安交通大学 (xjtu.edu.cn)](https://webvpn.xjtu.edu.cn/https/77726476706e69737468656265737421f1f940d23f3a7c45300d8db9d6562d/zh)
    - 校内访问 [社团首页 - 学生网络管理协会 - 西安交通大学 (xjtu.edu.cn)](http://ana.xjtu/edu.cn)

    复制积分信息第一行的密码，用于解压压缩包和作为<token>
3. 进入[xjtuana/ivpn: ivpn client (github.com)](https://github.com/xjtuana/ivpn)下拉，点击红色下划线部分
    > 如果不能正常访问，需要联系部长加入组织
    
    ![1](../picture/%E5%86%85%E7%BD%91%E8%AE%BF%E9%97%AE/%E5%86%85%E7%BD%91%E8%AE%BF%E9%97%AE-1.png)
    
    跳转后点击

    ![2](../picture/%E5%86%85%E7%BD%91%E8%AE%BF%E9%97%AE/%E5%86%85%E7%BD%91%E8%AE%BF%E9%97%AE-2.png)

    跳转后下拉，在红圈所属部分下载对应文件

    ![3](../picture/%E5%86%85%E7%BD%91%E8%AE%BF%E9%97%AE/%E5%86%85%E7%BD%91%E8%AE%BF%E9%97%AE-3.png)

    以amd64为例 ，  点击下载，在桌面上解压

    ![4](../picture/%E5%86%85%E7%BD%91%E8%AE%BF%E9%97%AE/%E5%86%85%E7%BD%91%E8%AE%BF%E9%97%AE-4.png)

4. 访问[Wintun – Layer 3 TUN Driver for Windows](https://www.wintun.net/),点击红色圈内按钮，下载后解压

    ![5](../picture/%E5%86%85%E7%BD%91%E8%AE%BF%E9%97%AE/%E5%86%85%E7%BD%91%E8%AE%BF%E9%97%AE-5.png)

    找到对应的.ddl文件，复制到跟.exe文件一个目录下

    ![6](../picture/%E5%86%85%E7%BD%91%E8%AE%BF%E9%97%AE/%E5%86%85%E7%BD%91%E8%AE%BF%E9%97%AE-6.png)
5. 在**管理员模式**下打开powershell
    使用如下命令启动程序（要保持该界面一直打开才能访问）
    ``` shell
    $  .\ivpnc.exe --device xjtuana --proxy wss://<token>@<addr>/ivp/ 
    ```
    > 校内<addr> 替换为ana.xjtu.edu.cn  
    > 校外<addr> 替换为xjtuana.com

    ![7](../picture/%E5%86%85%E7%BD%91%E8%AE%BF%E9%97%AE/%E5%86%85%E7%BD%91%E8%AE%BF%E9%97%AE-7.png)
6. 网页输入git.ana即可访问内网
    