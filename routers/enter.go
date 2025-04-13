package routers

import (
	"GoBlog/global"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	apiRouterGroup := r.Group("api")
	//系统配置api
	routerGroupApp := RouterGroup{apiRouterGroup}
	routerGroupApp.SettingsRouter()
	return r
}
