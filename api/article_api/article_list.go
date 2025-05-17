package article_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"GoBlog/utils/jwts"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (ArticleApi) ArticleListAdmin(c *gin.Context) {
	_cliams, _ := c.Get("claims")
	claims := _cliams.(*jwts.CustomClaims)

	var users models.UserModel
	err := global.DB.Take(&users, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("管理员用户不存在", c)
		return
	}

	var articles []models.ArticleModel
	result := global.DB.Find(&articles)
	if result.Error != nil {
		res.FailWithMsg("获取文章列表失败", c)
		return
	}

	// 提取所需字段
	var articleList []map[string]interface{}
	for _, article := range articles {
		articleData := map[string]interface{}{
			"id":             article.ID,
			"title":          article.Title,
			"abstract":       article.Abstract,
			"user_id":        article.UserID,
			"look_count":     article.LookCount,
			"comment_count":  article.CommentCount,
			"digg_count":     article.DiggCount,
			"collects_count": article.CollectsCount,
		}
		articleList = append(articleList, articleData)
	}

	res.OkWithData(articleList, c)
}

func (ArticleApi) GetArticle(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		res.FailWithMsg("无效的文章ID", c)
		return
	}
	//增加浏览量
	ArticleApi{}.ArticleUpdateLookCount(c)
	var article models.ArticleModel
	result := global.DB.Preload("TagModels").First(&article, id)
	if result.Error != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	res.OkWithData(article, c)
}
