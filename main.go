package main

import (
	"GoBlog/core"
	"GoBlog/global"
	"GoBlog/routers"
)

func main() {
	//	读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	//	gorm的连接
	global.DB = core.InitGorm()
	//初始化routers
	r := routers.InitRouter()
	global.Log.Info(global.Config.System.Addr())
	r.Run(global.Config.System.Addr())
}
