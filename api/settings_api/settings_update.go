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
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
	}
	fmt.Println("before", global.Config)
	global.Config.SiteInfo = cr
	fmt.Println("After", global.Config)
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OkWith(c)

}
