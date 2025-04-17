package images_api

import (
	"GoBlog/models"
	"GoBlog/models/res"
	"GoBlog/service/common"
	"github.com/gin-gonic/gin"
)

// ImageList 图片列表
// @Summary 图片列表
// @Description 图片列表
// @Tags 图片管理
// @Param data body models.PageInfo true "查询参数"
// @Router /api/images [get]
// @produce json
// @Accept json
// @Success 200 {object} res.Response{}
func (ImagesAPI) ImageList(c *gin.Context) {
	var page models.PageInfo
	err := c.ShouldBindQuery(&page)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}
	list, count, err := common.ComList(models.BannerModel{}, common.Option{
		page,
		false,
	})
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}
	res.OkWithList(list, count, c)
}
