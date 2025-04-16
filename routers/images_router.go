package routers

import "GoBlog/api"

func (r RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesAPI
	r.GET("images", imagesApi.ImageList)
	r.POST("images", imagesApi.ImageUploadView)
	r.DELETE("images", imagesApi.ImageRemove)
}
