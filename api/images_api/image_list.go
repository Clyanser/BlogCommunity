package images_api

import (
	"GoBlog/models"
	"GoBlog/models/res"
	"GoBlog/service/common"
	"github.com/gin-gonic/gin"
)

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
