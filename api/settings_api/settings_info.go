package settings_api

import (
	"GoBlog/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsAPI) SettingsInfo(c *gin.Context) {
	//c.JSON(200, gin.H{
	//	"message": "系统信息",
	//})
	res.Ok(map[string]string{}, "xxx", c)
}
