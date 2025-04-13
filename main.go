package main

import (
	"GoBlog/core"
	"GoBlog/global"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	//	读取配置文件
	core.InitConf()
	fmt.Println(global.Config)
	//初始化日志
	global.Log = core.InitLogger()
	global.Log.Warning("520")
	global.Log.Error("111")
	global.Log.Info("11111")

	logrus.Warning("520")
	logrus.Error("111")
	logrus.Info("11111")
	//	gorm的连接
	global.DB = core.InitGorm()
	fmt.Println(global.DB)
}
