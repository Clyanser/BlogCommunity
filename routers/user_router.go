package routers

import (
	"GoBlog/api"
	"GoBlog/middleware"
)

func (r RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserAPI
	r.POST("user", userApi.EmailLogin)
	r.GET("user", middleware.JwtAuth(), userApi.UserList)
	r.PUT("user_role", middleware.JwtAdmin(), userApi.UserUpdateRole)
	r.PUT("user_password", middleware.JwtAuth(), userApi.UserUpdatePassword)
	r.POST("user_logout", middleware.JwtAuth(), userApi.Logout)
	r.DELETE("user_delete", middleware.JwtAdmin(), userApi.UserRemove)
}
