# api-micro-mail

邮件微服务

### 使用方法

建议使用 `Postman` 等具有可视化的工具对服务进行调试。

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
go get git.ana/xjtuana/api-micro-mail
```

#### 命令

1. 将 `./config.example.toml` 复制一个副本 `./config.toml`；
2. 修改 `[api.micro.mail]` 中数据库地址，以及 `[api.micro.mail.smtp]` 中配置的 SMTP 邮件服务器及账号密码；
3. 运行 `go mod download` 下载依赖；
4. 运行 `go run cmd/server/main.go` 运行服务。

#### 容器

1. 按照如下方式组织 `docker-compose.yaml`；
2. 运行 `docker-compose up --build` 即可启动所有服务。

```yaml
version: "3.9"

services:
  api-micro-mail:
    image: registry.dev.xjtuana.cn/xjtuana/api-micro-mail:latest
    container_name: xjtuana-api-micro-mail
    depends_on:
      mariadb:
        condition: service_healthy
    restart: unless-stopped
    volumes:
      - type: bind
        source: ./config.toml # 指向本微服务目录下的配置文件
        target: /etc/xjtuana-api/config.toml
        read_only: true
    networks:
      br-xjtuana: {}
    expose:
      - 8080 # 仅暴露

  api-micro-other:
    build: . # 指向当前开发的微服务目录
    image: registry.dev.xjtuana.cn/xjtuana/api-micro-other:latest
    container_name: xjtuana-api-micro-other
    depends_on:
      mariadb:
        condition: service_healthy
    restart: unless-stopped
    volumes:
      - type: bind
        source: ./config.toml # 指向当前开发的微服务目录下的配置文件
        target: /etc/xjtuana-api/config.toml
        read_only: true
    networks:
      br-xjtuana: {}
    ports:
      - 8080:8080

  mariadb:
    image: ghcr.io/linuxserver/mariadb:latest
    container_name: xjtuana-mariadb
    healthcheck:
      test: ["CMD", "mysqladmin", "ping", "-h", "localhost"]
      timeout: 10s
      retries: 30
    restart: unless-stopped
    networks:
      br-xjtuana: {}
    expose:
      - 3306
    environment:
      - PUID=${MARIADB_PUID:-1000}
      - PGID=${MARIADB_PGID:-1000}
      - TZ=${MARIADB_TZ:-Asia/Shanghai}
      - MYSQL_ROOT_PASSWORD=${MARIADB_ROOT_PASSWORD:-password}
      - MYSQL_DATABASE=${MARIADB_DATABASE:-database}
      - MYSQL_USER=${MARIADB_USER:-user}
      - MYSQL_PASSWORD=${MARIADB_PASSWORD:-password}

networks:
  br-xjtuana:
    driver: bridge
    name: br-xjtuana
```
