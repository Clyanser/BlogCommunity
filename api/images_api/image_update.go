package images_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"github.com/gin-gonic/gin"
)

type ImageUpdate struct {
	ID   uint   `json:"id" binding:"required" msg:"请输入ID名称"`
	Name string `json:"name" binding:"required" msg:"请输入文件名称"`
}

func (ImagesAPI) ImageUpdate(c *gin.Context) {
	var cr ImageUpdate
	if err := c.ShouldBind(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var imageMode models.BannerModel
	if err := global.DB.Take(&imageMode, cr.ID).Error; err != nil {
		res.FailWithMsg("文件不存在", c)
		return
	}
	//	修改逻辑
	if err := global.DB.Model(&imageMode).Update("name", cr.Name).Error; err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OkWithMsg("修改成功", c)
	return
}
