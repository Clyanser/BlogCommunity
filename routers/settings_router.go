package routers

import (
	"GoBlog/api"
)

func (r RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsAPI
	//获取系统信息
	r.GET("settings/:uri", settingsApi.SettingsInfo)
	//修改系统信息
	r.PUT("settings/:uri", settingsApi.SettingsInfoUpdate)
	////获取邮箱信息
	//r.GET("settings_email", settingsApi.SettingsEmailInfo)
	////修改邮箱信息
	//r.PUT("settings_email", settingsApi.SettingsEmailUpdate)
}
