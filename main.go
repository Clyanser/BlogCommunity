package main

import (
	"GoBlog/core"
	"GoBlog/flag"
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
