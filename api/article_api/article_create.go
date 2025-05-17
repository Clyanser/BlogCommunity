package article_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/ctype"
	"GoBlog/models/res"
	"GoBlog/utils/jwts"
	"github.com/gin-gonic/gin"
)

type ArticleRequest struct {
	Title     string      `json:"title" binding:"required" msg:"请输入标题" structs:"title"`
	Abstract  string      `json:"abstract" binding:"required" msg:"请输入文章简介" structs:"abstract"`
	Content   string      `json:"content" binding:"required" msg:"请输入文章内容" structs:"content"`
	Category  string      `json:"category" binding:"required" msg:"请选择文章类别" structs:"category"`
	Source    string      `json:"source" structs:"source"`
	Link      string      `json:"link" structs:"link"`
	CoverPath string      `json:"coverPath" binding:"required" msg:"请选择文章封面" structs:"coverPath"`
	Tags      ctype.Array `json:"tags" msg:"请选择文章标签" structs:"tags"`                            // 标签ID列表
	BannerID  uint        `json:"bannerID" binding:"required" msg:"请输入封面ID" structs:"bannerID"` // 封面ID
}

func (ArticleApi) CreateArticle(c *gin.Context) {
	_cliams, _ := c.Get("claims")
	claims := _cliams.(*jwts.CustomClaims)

	var cr ArticleRequest
	//参数校验
	err := c.ShouldBind(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//检查用户是否存在
	var users models.UserModel
	err = global.DB.Take(&users, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	// 检查封面是否存在
	var banner models.BannerModel
	result := global.DB.First(&banner, cr.BannerID)
	if result.Error != nil {
		global.Log.Errorf("Banner with ID %d not found: %v", cr.BannerID, result.Error)
		res.FailWithMsg("封面不存在！", c)
		return
	}
	//检查tags数据是否在数据库存在
	// 检查标签是否存在
	var tagModels []models.TagModel
	for _, tagID := range cr.Tags {
		var tag models.TagModel
		result = global.DB.First(&tag, tagID)
		if result.Error != nil {
			global.Log.Errorf("Tag with ID %d not found: %v", tagID, result.Error)
			res.FailWithMsg("标签不存在！", c)
			return
		}
		tagModels = append(tagModels, tag)
	}
	//添加逻辑
	err = global.DB.Create(&models.ArticleModel{
		Title:     cr.Title,
		Abstract:  cr.Abstract,
		Content:   cr.Content,
		Category:  cr.Category,
		Source:    cr.Source,
		Link:      cr.Link,
		CoverPath: cr.CoverPath,
		UserID:    claims.UserID,
		NickName:  claims.Nickname,
		BannerID:  cr.BannerID,
		TagModels: tagModels, // 使用验证过的标签切片
		Tags:      cr.Tags,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("发布文章失败！", c)
		return
	}
	res.OkWithMsg("发布文章成功！", c)
}
