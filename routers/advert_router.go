package routers

import "GoBlog/api"

func (r RouterGroup) AdvertRouter() {
	advertApi := api.ApiGroupApp.AdvertAPI
	r.POST("adverts", advertApi.AdvertCreat)
}
