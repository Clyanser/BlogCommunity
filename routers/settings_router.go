package routers

import (
	"GoBlog/api"
)

func (r RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsAPI
	r.GET("settings", settingsApi.SettingsInfo)
}
