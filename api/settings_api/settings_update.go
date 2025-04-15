package settings_api

import (
	"GoBlog/config"
	"GoBlog/core"
	"GoBlog/global"
	"GoBlog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (SettingsAPI) SettingsInfoUpdate(c *gin.Context) {

	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	switch cr.Uri {
	case "site":
		var info config.SiteInfo
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ParamsError, c)
		}
		fmt.Println("before", global.Config)
		global.Config.SiteInfo = info
		fmt.Println("After", global.Config)
	case "email":
		var info config.Email
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ParamsError, c)
		}
		fmt.Println("before", global.Config)
		global.Config.Email = info
		fmt.Println("After", global.Config)
	case "qq":
		var info config.QQ
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ParamsError, c)
		}
		fmt.Println("before", global.Config)
		global.Config.QQ = info
		fmt.Println("After", global.Config)
	case "qiniu":
		var info config.QiNiu
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ParamsError, c)
		}
		fmt.Println("before", global.Config)
		global.Config.QiNiu = info
		fmt.Println("After", global.Config)
	case "jwt":
		var info config.Jwt
		err := c.ShouldBindJSON(&info)
		if err != nil {
			res.FailWithCode(res.ParamsError, c)
		}
		fmt.Println("before", global.Config)
		global.Config.Jwt = info
		fmt.Println("After", global.Config)
	default:
		res.FailWithMsg("没有对应的配置", c)
	}
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OkWith(c)

}
