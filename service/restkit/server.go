package restkit

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/cast"
	"mizuki/project/core-kit/service/configkit"
	"mizuki/project/core-kit/service/logkit"
	"mizuki/project/core-kit/service/restkit/middleware"
	router2 "mizuki/project/core-kit/service/restkit/router"
)

var router *router2.Router

//func defaultEngine2() *gin.Engine {
//	engine := gin.New()
//	//engine.Use(middleware.Log())
//	//engine.Use(middleware.Cors())
//	return engine
//}

func defaultEngine() {
	router = &router2.Router{
		IsGroup: false,
		Proxy:   iris.New(),
		Path:    "",
	}
	//router.Proxy.Use(recover2.New())
	router.Use(middleware.Recover())
	router.Use(middleware.Log())
	router.Use(middleware.Cors())
	// add base path
	base := configkit.GetStringD(ConfigKeyRestServerBase)
	if base != "" {
		if base[0] != '/' {
			base = "/" + base
		}
		router.ProxyGroup = router.Proxy.Party(base)
		router.IsGroup = true
		//router.Path = base
	}
}

func Run() {
	if router == nil {
		defaultEngine()
	}
	port := cast.ToString(configkit.Get(ConfigKeyRestServerPort, 8080))
	logkit.Info("Listening and serving HTTP on " + port)
	//err := http.ListenAndServe(":" + port, middleware.Session.LoadAndSave(router))
	router.RegisterSwagger()
	err := router.Proxy.Run(iris.Addr(":" + port))
	if err != nil {
		logkit.Fatal(err.Error())
	}
}

// 导入业务模块，其中的路由和中间件
func AddMod(modInit func(r *router2.Router)) {
	if router == nil {
		defaultEngine()
	}
	modInit(router)
}
