package bootstrap

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"time"
	"web_iris/golang_mall/conf"
)

type Configurator func(bootstrapper *Bootstrapper) //定义引导方法

type Bootstrapper struct {
	//引导方法所需类型
	*iris.Application
	AppName     string    //app名称
	AppOwner    string    // App所有者
	AppSpawDate time.Time //更新时间
}

//实例化bootstrapper 初始iris app
func New(appName, appOwner string, cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		Application: iris.New(),
		AppName:     appName,
		AppOwner:    appOwner,
		AppSpawDate: time.Now(),
	}

	for _, cfg := range cfgs {
		cfg(b) //给这个空的引导方法附上地址
	}
	return b
}

//初始views模板
func (b *Bootstrapper) SetupViews(viewsDir string) {
	htmlEngine := iris.HTML(viewsDir, ".html") //layout 是布局文件
	htmlEngine.Reload(true)                    //上线时关闭,每一次都会加载html
	//把一个unix时间直接就可以调用转换成字符串
	htmlEngine.AddFunc("FromUnixtime", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeform)
	})
	htmlEngine.AddFunc("FromUnixtimeShort", func(t int) string {
		dt := time.Unix(int64(t), int64(0))
		return dt.Format(conf.SysTimeformShort)
	})
	b.RegisterView(htmlEngine)
}

//处理报错
func (b *Bootstrapper) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(cxt iris.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  cxt.GetStatusCode(),
			"message": cxt.Values().GetString("message"),
		}
		if jsonOutput := cxt.URLParamExists("json"); jsonOutput { //判断 包含JS
			cxt.JSON(err)
			return
		}
		cxt.ViewData("err", err)
		cxt.ViewData("Title", "Error")
		//cxt.View("shared/error.html")
	})
}

//添加配置  方法

func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

func (b *Bootstrapper) setupCron() { //执行方法
	//cron.ConfigueAppOneCron()
	//comm.IsExist("banner")

}

const (
	StaticAssets = "./public/" //静态文件
	//Favicon      = "favicon.ico" //网站图标
)

func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.SetupViews("./views/") //html文件目录
	b.SetupErrorHandlers()   //模板报错任务
	//b.Favicon(StaticAssets, Favicon)                               //网站图标
	b.StaticWeb(StaticAssets[1:len(StaticAssets)-1], StaticAssets) //网站站点
	//b.setupCron()                                                  //启动计划任务
	b.Use(recover.New()) //异常时-日志
	b.Use(logger.New())  //日志处理-日志
	return b
}

func (b *Bootstrapper) Listen(addr string, cfs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfs...)
}
