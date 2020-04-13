package main

import (
	"go-web/boot"
	"go-web/controllers"
)

func main() {
	// 注册路由
	controllers.RegisteRouter(boot.App.Engine)

	boot.App.Serve()
}
