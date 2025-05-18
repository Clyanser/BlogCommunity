package routers

import (
	"GoBlog/api"
	"GoBlog/middleware"
)

func (r RouterGroup) ArticleCollectRouter() {
	articleCollectApi := api.ApiGroupApp.ArticleCollectApi
	r.POST("articles/:id/collect", middleware.JwtAuth(), articleCollectApi.CollectArticle)
	r.DELETE("articles/:id/collect", middleware.JwtAuth(), articleCollectApi.UncollectArticle)
}
