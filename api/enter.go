package api

import (
	"GoBlog/api/advert_api"
	"GoBlog/api/article_api"
	"GoBlog/api/images_api"
	"GoBlog/api/menu_api"
	"GoBlog/api/message_api"
	"GoBlog/api/settings_api"
	"GoBlog/api/tag_api"
	"GoBlog/api/user_api"
)

type ApiGroup struct {
	SettingsAPI settings_api.SettingsAPI
	ImagesAPI   images_api.ImagesAPI
	AdvertAPI   advert_api.AdvertApi
	MenuAPI     menu_api.MenuAPI
	UserAPI     user_api.UserApi
	TagAPI      tag_api.TagApi
	MessageAPI  message_api.MessageAPI
	ArticleAPI  article_api.ArticleApi
}

// 实例化对象
var ApiGroupApp = new(ApiGroup)
