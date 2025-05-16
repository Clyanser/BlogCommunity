package routers

import (
	"GoBlog/api"
	"GoBlog/middleware"
)

func (r RouterGroup) MessageRouter() {
	messageApi := api.ApiGroupApp.MessageAPI
	r.POST("messages", middleware.JwtAuth(), messageApi.MessageCreate)
	r.GET("messages", middleware.JwtAdmin(), messageApi.MessageAdminList)
	r.GET("messages_user", middleware.JwtAuth(), messageApi.MessageUserList)

}
