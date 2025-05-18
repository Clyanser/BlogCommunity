package article_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"GoBlog/utils/jwts"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func (ArticleApi) DeleteArticle(c *gin.Context) {
	//获取文章id
	articleIDStr := c.Param("id")
	articleID, err := strconv.Atoi(articleIDStr)
	if err != nil {
		res.FailWithMsg("无效的文章ID", c)
		return
	}

	// 获取当前登录用户的ID
	_cliams, _ := c.Get("claims")
	claims := _cliams.(*jwts.CustomClaims)

	var article models.ArticleModel
	result := global.DB.Preload("CommentModels").First(&article, articleID).Error
	if result != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	// 检查是否是文章的作者
	if article.UserID != claims.UserID {
		res.FailWithMsg("无权限删除此文章", c)
		return
	}

	// 开始数据库事务
	tx := global.DB.Begin()
	if tx.Error != nil {
		res.FailWithMsg("开始事务失败", c)
		return
	}

	// 递归删除所有评论及其子评论
	for _, comment := range article.CommentModels {
		err = deleteSubComments(tx, &comment)
		if err != nil {
			tx.Rollback()
			global.Log.Error(err)
			res.FailWithMsg("删除评论失败", c)
			return
		}
		// 删除顶级评论
		result = tx.Delete(&comment).Error
		if result != nil {
			tx.Rollback()
			global.Log.Error(result)
			res.FailWithMsg("删除评论失败", c)
			return
		}
	}

	// 删除文章与标签的关系
	result = tx.Table("article_tag_models").Where("article_model_id = ?", articleID).Delete(nil).Error
	if result != nil {
		tx.Rollback()
		global.Log.Error(result)
		res.FailWithMsg("删除文章标签关系失败", c)
		return
	}

	// 删除文章
	result = tx.Delete(&article).Error
	if result != nil {
		tx.Rollback()
		global.Log.Error(result)
		res.FailWithMsg("删除文章失败", c)
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		res.FailWithMsg("提交事务失败", c)
		return
	}

	res.OkWithMsg("文章删除成功", c)
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
