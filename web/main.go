package main

import (
	"fmt"

	"web_iris/golang_mall/bootstrap"
	"web_iris/golang_mall/web/routes"

	"web_iris/golang_mall/web/middlewares/identity"
)

var port = 8080

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Go抽奖系统", "杰sir")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure) //加配置
	return app
}

func main() {
	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))
}
