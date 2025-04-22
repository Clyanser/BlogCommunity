package routers

import "GoBlog/api"

func (r RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserAPI
	r.GET("user", userApi.UserList)
	r.POST("user", userApi.EmailLogin)

}
