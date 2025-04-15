package settings_api

import (
	"GoBlog/global"
	"GoBlog/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsAPI) SettingsInfo(c *gin.Context) {
	res.OkWithData(global.Config.SiteInfo, c)
}
