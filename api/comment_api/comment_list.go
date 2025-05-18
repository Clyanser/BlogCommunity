package comment_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (CommentApi) GetCommentList(c *gin.Context) {
	// 获取文章ID
	articleIDStr := c.Param("id")
	articleID, err := strconv.Atoi(articleIDStr)
	if err != nil {
		res.FailWithMsg("无效的文章ID", c)
		return
	}
	//获取评论列表
	var comments []models.CommentModel
	result := global.DB.Where("article_id = ? AND parent_comment_id IS NULL", articleID).Find(&comments).Error
	if result != nil {
		res.FailWithMsg("获取评论失败", c)
		return
	}
	for i := range comments {
		loadSubComments(&comments[i])
		global.DB.Model(&comments[i]).Association("User").Find(&comments[i].User)
	}

	res.OkWithData(comments, c)

}

func loadSubComments(comment *models.CommentModel) {
	//查询子评论列表
	var subComments []models.CommentModel
	result := global.DB.Where("parent_comment_id = ?", comment.ID).Find(&subComments).Error
	if result != nil {
		return
	}

	// 将 []models.CommentModel 转换为 []*models.CommentModel
	comment.SubComments = make([]*models.CommentModel, len(subComments))
	for i := range subComments {
		comment.SubComments[i] = &subComments[i]
	}

	// 递归加载子评论和用户信息
	for _, subComment := range comment.SubComments {
		loadSubComments(subComment)
		global.DB.Model(subComment).Association("User").Find(&subComment.User)
	}
}
