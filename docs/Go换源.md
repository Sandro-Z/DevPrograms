# Go 换源方法

## 为什么需要换源？

通常来说，我们在终端运行`go env`以后，可以看到输出内容中有一行

```
GOPROXY="https://proxy.golang.org,direct"
```

我们在国内是无法访问这个地址的，也就意味着当你使用`go mod tidy`来整理模块的时候，会报timeout。

## 解决方法

### 第一种，给Go换源（推荐，全平台通用）

在终端运行

```
go env -w GOPROXY=https://goproxy.cn,direct
go env | grep GOPROXY                      # 如果你不是Linux用户，运行go env，然后找到GOPROXY那一行
```

如果看见如下输出

```
GOPROXY="https://goproxy.cn,direct"
```

说明Go语言换源成功，之后使用`go mod tidy`就可以自动导入所有包

### 第二种，使用魔法（不推荐，Linux only）

如果你自己有socks5代理，可以使用如下命令，在终端导入代理地址，使得你可以访问Go官方源。注意，该导入命令是一次性的，仅当前终端有效。

```
export HTTP_PROXY="YOUR PROXY ADDRESS"    # Example: export HTTP_PROXY="socks5://127.0.0.1:1080"
export HTTPS_PROXY="YOUR PROXY ADDRESS"
export | grep HTTP
```

如果看见如下输出

```
declare -x HTTPS_PROXY="YOUR PROXY ADDRESS"
declare -x HTTP_PROXY="YOUR PROXY ADDRESS"
```

说明导入成功，可以使用Go官方源了。

**我不想每次打开终端都输入一遍怎么办？**

在终端输入如下命令，请直接把一整行命令粘贴到终端（终端粘贴需要使用Ctrl+Shift+v）并输入回车

```
echo "# http/https proxy"$'\n'"export HTTP_PROXY=\"YOUR PROXY ADDRESS\""$'\n'"export HTTPS_PROXY=\"YOUR PROXY ADDRESS\"" >> ~/.bashrc
```

重启终端，`export | grep HTTP`，看见如下内容，说明成功。

```
declare -x HTTPS_PROXY="YOUR PROXY ADDRESS"
declare -x HTTP_PROXY="YOUR PROXY ADDRESS"
```





