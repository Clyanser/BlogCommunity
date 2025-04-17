package routers

import (
	"GoBlog/api"
)

func (r RouterGroup) MenuRouter() {
	menuRouter := api.ApiGroupApp.MenuAPI
	r.POST("menu", menuRouter.MenuCreate)
}
