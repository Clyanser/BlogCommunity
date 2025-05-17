package routers

import (
	"GoBlog/api"
	"GoBlog/middleware"
)

func (r RouterGroup) ArticleRouter() {
	articleApi := api.ApiGroupApp.ArticleAPI
	r.POST("/articles", middleware.JwtAuth(), articleApi.CreateArticle)
	r.GET("/articles", middleware.JwtAdmin(), articleApi.ArticleListAdmin)
	r.GET("/articles_details/:id", middleware.JwtAuth(), articleApi.GetArticle)
	r.PUT("/articles/:id", middleware.JwtAuth(), articleApi.ArticleUpdate)
	r.POST("/articles_digg/:id", middleware.JwtAuth(), articleApi.ArticleUpdateDiggCount)
}
