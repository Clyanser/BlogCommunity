package comment_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"GoBlog/utils/jwts"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (CommentApi) DeleteComment(c *gin.Context) {
	_cliams, _ := c.Get("claims")
	claims := _cliams.(*jwts.CustomClaims)

	commentIDStr := c.Param("id")
	commentID, err := strconv.Atoi(commentIDStr)
	if err != nil {
		res.FailWithMsg("无效的评论ID", c)
		return
	}

	var comment models.CommentModel
	result := global.DB.Preload("SubComments").First(&comment, commentID).Error
	if result != nil {
		res.FailWithMsg("评论不存在", c)
		return
	}

	// 检查是否是评论的作者
	if comment.UserID != claims.UserID {
		res.FailWithMsg("无权限删除此评论", c)
		return
	}

	// 开始数据库事务
	tx := global.DB.Begin()
	if tx.Error != nil {
		res.FailWithMsg("开始事务失败", c)
		return
	}

	// 递归删除所有子评论
	err = deleteSubComments(tx, &comment)
	if err != nil {
		tx.Rollback()
		global.Log.Error(err)
		res.FailWithMsg("删除子评论失败", c)
		return
	}

	// 删除指定的评论
	result = tx.Delete(&comment).Error
	if result != nil {
		tx.Rollback()
		global.Log.Error(result)
		res.FailWithMsg("删除评论失败", c)
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		res.FailWithMsg("提交事务失败", c)
		return
	}

	res.OkWithMsg("评论删除成功", c)
}

func deleteSubComments(tx *gorm.DB, comment *models.CommentModel) error {
	for _, subComment := range comment.SubComments {
		// 递归删除子评论的子评论
		err := deleteSubComments(tx, subComment)
		if err != nil {
			return err
		}
		// 删除子评论
		result := tx.Delete(subComment).Error
		if result != nil {
			return result
		}
	}
	return nil
}
