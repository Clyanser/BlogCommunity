package routers

import (
	"GoBlog/api"
	"GoBlog/middleware"
)

func (r RouterGroup) MessageRouter() {
	messageApi := api.ApiGroupApp.MessageAPI
	r.POST("messages", middleware.JwtAuth(), messageApi.MessageCreate)

}
