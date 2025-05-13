package routers

import "GoBlog/api"

func (r RouterGroup) TagRouter() {
	tagApi := api.ApiGroupApp.TagAPI
	r.GET("tags", tagApi.TagList)
	r.POST("tags", tagApi.TagCreate)
	r.PUT("tags/:id", tagApi.TagUpdate)
	r.DELETE("tags", tagApi.TagRemove)
}
