package routers

import "GoBlog/api"

func (r RouterGroup) AdvertRouter() {
	advertApi := api.ApiGroupApp.AdvertAPI
	r.GET("adverts", advertApi.AdvertList)
	r.POST("adverts", advertApi.AdvertCreat)
	r.PUT("adverts/:id", advertApi.AdvertUpdate)
	r.DELETE("adverts", advertApi.AdvertRemove)
}
