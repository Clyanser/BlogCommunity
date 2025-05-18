package routers

import (
	"GoBlog/api/comment_api"
	"GoBlog/middleware"
)

func (r RouterGroup) CommentRouter() {
	commentApi := comment_api.CommentApi{}
	r.POST("/comments/:id", middleware.JwtAuth(), commentApi.CommentCreate)
	r.GET("/comments/:id", commentApi.GetCommentList)
	r.POST("/comments_digg/:id", commentApi.UpdateDiggCount)
	r.DELETE("/comments/:id", middleware.JwtAuth(), commentApi.DeleteComment)
}
