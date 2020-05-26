package identity

import "github.com/kataras/iris"

func before(ctx iris.Context) {
	shareInformation := "这是处理程序之间可共享的信息"

	requestPath := ctx.Path()
	println("mainHandler之前: " + requestPath)

	ctx.Values().Set("info", shareInformation)
	ctx.Next() // 执行下一个处理器。
}

func after(ctx iris.Context) {
	println("mainHandler后")
}

func mainHandler(ctx iris.Context) {
	println("内部mainHandler")

	// 获取 "before" 处理器中的设置的 "info" 值。
	info := ctx.Values().GetString("info")

	// 响应客户端
	ctx.HTML("<h1>响应</h1>")
	ctx.HTML("<br/> Info: " + info)

	ctx.Next() // execute the "after".
}
