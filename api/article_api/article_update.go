package article_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"GoBlog/utils/jwts"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (ArticleApi) ArticleUpdate(c *gin.Context) {
	_cliams, _ := c.Get("claims")
	claims := _cliams.(*jwts.CustomClaims)

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.FailWithMsg("无效的文章ID", c)
		return
	}
	//检查用户是否存在
	var users models.UserModel
	err = global.DB.Take(&users, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}

	var cr ArticleRequest
	//参数校验
	err = c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//校验文章是否存在
	var article models.ArticleModel
	result := global.DB.First(&article, id)
	if result.Error != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}
	//更新文章逻辑
	article.Title = cr.Title
	article.Abstract = cr.Abstract
	article.Content = cr.Content
	article.Category = cr.Category
	article.Source = cr.Source
	article.Link = cr.Link
	article.CoverPath = cr.CoverPath
	article.BannerID = cr.BannerID

	//更新标签
	// 更新标签
	var tagModels []models.TagModel
	for _, tagID := range cr.Tags {
		var tag models.TagModel
		result = global.DB.First(&tag, tagID)
		if result.Error != nil {
			res.FailWithMsg("标签不存在！", c)
			return
		}
		tagModels = append(tagModels, tag)
	}
	article.TagModels = tagModels
	// 保存更新后的文章
	if err := global.DB.Save(&article).Error; err != nil {
		res.FailWithMsg("更新文章失败", c)
		return
	}

	res.OkWithMsg("更新文章成功", c)
}

func (ArticleApi) ArticleUpdateLookCount(c *gin.Context) {
	ArticleApi{}.updateCounterField(c, "look_count")
}

func (ArticleApi) ArticleUpdateCommentCount(c *gin.Context) {
	ArticleApi{}.updateCounterField(c, "comment_count")
}

func (ArticleApi) ArticleUpdateDiggCount(c *gin.Context) {
	ArticleApi{}.updateCounterField(c, "digg_count")
}

func (ArticleApi) ArticleUpdateCollectsCount(c *gin.Context) {
	ArticleApi{}.updateCounterField(c, "collects_count")
}

func (ArticleApi) updateCounterField(c *gin.Context, fieldName string) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.FailWithMsg("无效的文章ID", c)
		return
	}

	var article models.ArticleModel
	result := global.DB.First(&article, id)
	if result.Error != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	switch fieldName {
	case "look_count":
		article.LookCount++
	case "comment_count":
		article.CommentCount++
	case "digg_count":
		article.DiggCount++
	case "collects_count":
		article.CollectsCount++
	default:
		res.FailWithMsg("无效的计数器字段", c)
		return
	}

	if err := global.DB.Save(&article).Error; err != nil {
		res.FailWithMsg("更新计数器失败", c)
		return
	}

	res.OkWithMsg("更新计数器成功", c)
}
