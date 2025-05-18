package comment_api

import (
	"GoBlog/api/article_api"
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"GoBlog/utils/jwts"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CommentRequest struct {
	Content         string `json:"content" binding:"required"`
	ParentCommentID *uint  `json:"parent_comment_id"` // 可选，用于子评论
}

func (CommentApi) CommentCreate(c *gin.Context) {
	_cliams, _ := c.Get("claims")
	claims := _cliams.(*jwts.CustomClaims)

	// 获取文章ID
	articleIDStr := c.Param("id")
	articleID, err := strconv.Atoi(articleIDStr)
	if err != nil {
		res.FailWithMsg("无效的文章ID", c)
		return
	}
	//参数绑定
	var cr CommentRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithMsg("参数绑定失败", c)
		return
	}

	//验证用户合法性
	var users models.UserModel
	err = global.DB.Take(&users, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}

	// 创建评论
	newComment := models.CommentModel{
		Content:         cr.Content,
		ArticleID:       uint(articleID),
		UserID:          claims.UserID,
		ParentCommentID: cr.ParentCommentID,
	}
	err = global.DB.Create(&newComment).Error
	if err != nil {
		global.Log.Errorf("创建评论失败 %s", err.Error())
		res.FailWithMsg("创建评论失败", c)
		return
	}
	article_api.ArticleApi{}.ArticleUpdateCommentCount(c)
	//增加父级评论评论数
	if cr.ParentCommentID != nil {
		article_api.ArticleApi{}.ArticleUpdateCommentCount(c)
		CommentApi{}.UpdateCommentCount(c, cr.ParentCommentID)
	}

	res.OkWithMsg("创建评论成功！", c)
}
