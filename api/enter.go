package api

import "GoBlog/api/settings_api"

type ApiGroup struct {
	SettingsAPI settings_api.SettingsAPI
}

// 实例化对象
var ApiGroupApp = new(ApiGroup)
