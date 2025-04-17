package images_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"github.com/gin-gonic/gin"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`                //图片路径
	Name string `gorm:"size:38" json:"name"` //图片名称
}

// @Summary 图片名称列表
// @Description 图片名称列表
// @Tags 图片管理
// @Param data body models.PageInfo true "查询参数"
// @Router /api/images_name [get]
// @produce json
// @Success 200 {object} res.Response{data=ImageResponse}
func (ImagesAPI) ImageNameList(c *gin.Context) {
	var imagesList []ImageResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imagesList)
	res.OkWithData(imagesList, c)
}
