# go-web 

## 一个简单的使用gin构建web api的工程实践例子

## 部署

**编译**
```
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o go-web main.go
```
**配置文件**

**服务器部署**


## 启动服务
```
CONF_PATH="<配置文件路径>" go run main.go
```