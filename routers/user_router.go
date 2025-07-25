package routers

import (
	"GoBlog/api"
	"GoBlog/middleware"
)

//var store = cookie.NewStore([]byte("jaory1019"))

func (r RouterGroup) UserRouter() {
	//r.Use(sessions.Sessions("sessionID", store))
	userApi := api.ApiGroupApp.UserAPI
	r.POST("user", userApi.EmailLogin)
	r.POST("user_create", middleware.JwtAdmin(), userApi.UserCreate)
	r.GET("user", middleware.JwtAuth(), userApi.UserList)
	r.PUT("user_role", middleware.JwtAdmin(), userApi.UserUpdateRole)
	r.PUT("user_password", middleware.JwtAuth(), userApi.UserUpdatePassword)
	r.POST("user_logout", middleware.JwtAuth(), userApi.Logout)
	r.DELETE("user_delete", middleware.JwtAdmin(), userApi.UserRemove)
	r.POST("user_bind_email", middleware.JwtAuth(), userApi.UserBindEmail)
}
