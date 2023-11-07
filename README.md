## 简介

ZeroSSL 提供免费的IP证书，并且可以支持文件验证。

本项目提供一个简易的认证服务器。

## 如何使用

1. 将认证txt文件放到项目根目录
2. 运行程序（release或者自己build均可）
3. 访问认证网站，测试是否能获取到文件内容

## build

```
go env -w GOOS=linux
go env -w GOARCH=amd64
go build -o ./build/zerossl_verify main.go
```