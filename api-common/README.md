# api-common

### 使用方法

#### 环境配置

`~/.gitconfig` 添加 `URL`

```
[url "ssh://git@git.ana/"]
    insteadOf = http://git.ana/
```

`go env` 添加环境变量

```bash
go env -w GOINSECURE=git.ana
go env -w GOPRIVATE=git.ana
```

`go get` 添加私有仓库

```bash
go get -u git.ana/xjtuana/api-common
```
