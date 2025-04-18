package routers

import (
	"GoBlog/api"
)

func (r RouterGroup) MenuRouter() {
	menuRouter := api.ApiGroupApp.MenuAPI
	r.POST("menu", menuRouter.MenuCreate)
	r.GET("menu_details/:id", menuRouter.MenuDetails)
	r.GET("menu", menuRouter.MenuList)
	r.GET("menu_names", menuRouter.MenuNameList)
	r.PUT("menu/:id", menuRouter.MenuUpdate)
	r.DELETE("menu", menuRouter.MenuDelete)
}
