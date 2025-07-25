package advert_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"github.com/gin-gonic/gin"
)

type AdvertReq struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`        //显示的标题
	Href   string `json:"href" binding:"required,url" msg:"非法的跳转链接" structs:"href"`    //跳转链接
	Images string `json:"images" binding:"required,url" msg:"图片地址非法" structs:"images"` //图片
	IsShow bool   `json:"is_show" msg:"请选择标题是否展示" structs:"is_show"`                   //是否展示
}

// AdvertCreat 添加广告
// @Summary 创建广告
// @Description 创建广告
// @Tags 广告管理
// @Param data body AdvertReq  true "表示多个参数"
// @Router /api/adverts [post]
// produce json
// @Success 200 {object} res.Response{}
func (AdvertApi) AdvertCreat(c *gin.Context) {
	var cr AdvertReq
	//参数校验
	err := c.ShouldBind(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//重复的判断
	var advert models.AdvertModel
	err = global.DB.Take(&advert, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMsg("该广告已存在", c)
		return
	}
	//添加逻辑
	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("添加广告失败！", c)
		return
	}
	res.OkWithMsg("添加广告成功！", c)
}
