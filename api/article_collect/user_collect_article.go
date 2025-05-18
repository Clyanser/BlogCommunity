package article_collect

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"GoBlog/utils/jwts"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (ArticleCollectApi) CollectArticle(c *gin.Context) {
	articleIDStr := c.Param("id")
	articleID, err := strconv.Atoi(articleIDStr)
	if err != nil {
		res.FailWithMsg("无效的文章ID", c)
		return
	}

	// 获取当前登录用户的ID
	_cliams, _ := c.Get("claims")
	claims := _cliams.(*jwts.CustomClaims)

	// 检查文章是否存在
	var article models.ArticleModel
	result := global.DB.First(&article, articleID).Error
	if result != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	// 检查用户是否已经收藏了该文章
	var userCollect models.UserCollectModel
	result = global.DB.Where("user_id = ? AND article_id = ?", claims.UserID, articleID).First(&userCollect).Error
	if result == nil {
		res.FailWithMsg("您已经收藏了这篇文章", c)
		return
	}

	// 创建新的收藏记录
	newCollect := models.UserCollectModel{
		UserID:    claims.UserID,
		ArticleID: uint(articleID),
	}

	result = global.DB.Create(&newCollect).Error
	if result != nil {
		global.Log.Errorf("创建收藏记录失败: %v", result)
		res.FailWithMsg("收藏文章失败", c)
		return
	}

	// 更新文章的收藏量
	article.CollectsCount++
	result = global.DB.Save(&article).Error
	if result != nil {
		global.Log.Errorf("更新文章收藏量失败: %v", result)
		res.FailWithMsg("更新文章收藏量失败", c)
		return
	}

	res.OkWithMsg("文章收藏成功", c)
}

// UncollectArticle 处理用户取消收藏文章的请求
func (ArticleCollectApi) UncollectArticle(c *gin.Context) {
	articleIDStr := c.Param("id")
	articleID, err := strconv.Atoi(articleIDStr)
	if err != nil {
		res.FailWithMsg("无效的文章ID", c)
		return
	}

	// 获取当前登录用户的ID
	_cliams, _ := c.Get("claims")
	claims := _cliams.(*jwts.CustomClaims)

	// 检查文章是否存在
	var article models.ArticleModel
	result := global.DB.First(&article, articleID).Error
	if result != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	// 检查用户是否已经收藏了该文章
	var userCollect models.UserCollectModel
	result = global.DB.Where("user_id = ? AND article_id = ?", claims.UserID, articleID).First(&userCollect).Error
	if result != nil {
		res.FailWithMsg("您还没有收藏这篇文章", c)
		return
	}

	// 删除收藏记录
	result = global.DB.Delete(&userCollect).Error
	if result != nil {
		global.Log.Errorf("删除收藏记录失败: %v", result)
		res.FailWithMsg("取消收藏文章失败", c)
		return
	}

	// 更新文章的收藏量
	if article.CollectsCount > 0 {
		article.CollectsCount--
		result = global.DB.Save(&article).Error
		if result != nil {
			global.Log.Errorf("更新文章收藏量失败: %v", result)
			res.FailWithMsg("更新文章收藏量失败", c)
			return
		}
	}

	res.OkWithMsg("文章取消收藏成功", c)
}
