package api

import (
	"GoBlog/api/advert_api"
	"GoBlog/api/images_api"
	"GoBlog/api/settings_api"
)

type ApiGroup struct {
	SettingsAPI settings_api.SettingsAPI
	ImagesAPI   images_api.ImagesAPI
	AdvertAPI   advert_api.AdvertApi
}

// 实例化对象
var ApiGroupApp = new(ApiGroup)
