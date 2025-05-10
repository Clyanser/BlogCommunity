package routers

import (
	"GoBlog/api"
	"GoBlog/middleware"
)

func (r RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserAPI
	r.GET("user", middleware.JwtAuth(), userApi.UserList)
	r.POST("user", userApi.EmailLogin)
	r.PUT("user_role", userApi.UserUpdateRole)
	r.PUT("user_password", middleware.JwtAuth(), userApi.UserUpdatePassword)

}
