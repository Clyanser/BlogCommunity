package routers

import (
	"GoBlog/api"
)

func (r RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsAPI
	//获取系统信息
	r.GET("settings", settingsApi.SettingsInfo)
	//修改系统信息
	r.PUT("settings", settingsApi.SettingsInfoUpdate)
}
