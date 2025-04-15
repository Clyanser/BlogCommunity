package routers

import "GoBlog/api"

func (r RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesAPI
	r.POST("images", imagesApi.ImageUploadView)
}
