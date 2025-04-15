package settings_api

import (
	"GoBlog/global"
	"GoBlog/models/res"
	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Uri string `uri:"uri"`
}

func (SettingsAPI) SettingsInfo(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}
	switch cr.Uri {
	case "site":
		res.OkWithData(global.Config.SiteInfo, c)
	case "email":
		res.OkWithData(global.Config.Email, c)
	case "qq":
		res.OkWithData(global.Config.QQ, c)
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
	default:
		res.FailWithMsg("没有对应的配置", c)
	}
}
