package settings_api

import (
	"GoBlog/config"
	"GoBlog/core"
	"GoBlog/global"
	"GoBlog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (SettingsAPI) SettingsEmailUpdate(c *gin.Context) {
	var cr config.Email
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
	}
	fmt.Println("before", global.Config.Email)
	global.Config.Email = cr
	fmt.Println("After", global.Config.Email)
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OkWith(c)
}
