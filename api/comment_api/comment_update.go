package comment_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (CommentApi) CommentUpdate(c *gin.Context) {

}

func (CommentApi) UpdateCommentCount(c *gin.Context, id *uint) {
	var comment models.CommentModel
	result := global.DB.First(&comment, id)
	if result.Error != nil {
		res.FailWithMsg("父级评论不存在", c)
		return
	}
	//增加逻辑
	comment.CommentCount++
	if err := global.DB.Save(&comment).Error; err != nil {
		res.FailWithMsg("更新计数器失败", c)
		return
	}
	res.OkWithMsg("更新计数器成功", c)
}

func (CommentApi) UpdateDiggCount(c *gin.Context) {
	CommentApi{}.updateCounterField(c, "digg_count")
}

func (CommentApi) updateCounterField(c *gin.Context, fieldName string) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.FailWithMsg("无效的父级评论ID", c)
		return
	}

	var comment models.CommentModel
	result := global.DB.First(&comment, id)
	if result.Error != nil {
		res.FailWithMsg("父级评论不存在", c)
		return
	}

	switch fieldName {
	case "digg_count":
		comment.DiggCount++
	default:
		res.FailWithMsg("无效的计数器字段", c)
		return
	}

	if err := global.DB.Save(&comment).Error; err != nil {
		res.FailWithMsg("更新计数器失败", c)
		return
	}

	res.OkWithMsg("更新计数器成功", c)
}
