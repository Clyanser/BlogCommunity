package main

import (
	"GoBlog/core"
	_ "GoBlog/docs"
	"GoBlog/flag"
	"GoBlog/global"
	"GoBlog/routers"
)

// @title GoBlog_server API文档
// @version 1.0
// @description API文档
// @host 127.0.0.1:9090
// @BasePath /
func main() {
	//	读取配置文件
	core.InitConf()
	//初始化日志
	global.Log = core.InitLogger()
	//	gorm的连接
	global.DB = core.InitGorm()
	//命令行参数绑定
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	//初始化routers
	r := routers.InitRouter()
	global.Log.Info(global.Config.System.Addr())
	err := r.Run(global.Config.System.Addr())
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
