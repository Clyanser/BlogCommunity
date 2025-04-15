package api

import (
	"GoBlog/api/images_api"
	"GoBlog/api/settings_api"
)

type ApiGroup struct {
	SettingsAPI settings_api.SettingsAPI
	ImagesAPI   images_api.ImagesAPI
}

// 实例化对象
var ApiGroupApp = new(ApiGroup)
