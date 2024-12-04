# 带有tagger的proto文件说明

## 使用的开源项目
[protoc-gen-gotag](https://github.com/srikrsna/protoc-gen-gotag/)

## tagger的作用
protobuf的tagger插件可以给protobuf生成的结构添加其他tag，如json tag，以便在dto中复用。

## 编译带有tagger的pb文件的方法

1. 安装protoc-gen-gotag项目：`go install github.com/srikrsna/protoc-gen-gotag@latest`

2. 定位protoc-gen-gotag包位置，通常为`$GOPATH/pkg/mod/github.com/srikrsna/protoc-gen-gotag@版本号`。

3. 构建protobuf对应的golang文件：`protoc -I $GOPATH/pkg/mod/github.com/srikrsna/protoc-gen-gotag@v0.6.2 -I . --go_out=:. user.proto`

4. 对golang文件添加tag：`protoc -I $GOPATH/pkg/mod/github.com/srikrsna/protoc-gen-gotag@版本号 -I . --plugin=protoc-gen-gotag=$GOPATH/bin/protoc-gen-gotag --gotag_out=:. user.proto`