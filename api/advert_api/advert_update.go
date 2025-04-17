package advert_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// AdvertUpdate 更新广告
// @Summary 更新广告
// @Description 更新广告
// @Tags 广告管理
// @Param data body AdvertReq true "广告的一些参数"
// @Router /api/adverts/:id [put]
// @produce json
// @Accept json
// @Success 200 {object} res.Response{data=string}
func (AdvertApi) AdvertUpdate(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertReq
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMsg("广告不存在！", c)
		return
	}
	//结构体转map的第三方包
	maps := structs.Map(&cr)
	err = global.DB.Model(&advert).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改广告失败", c)
		return
	}
	res.FailWithMsg("修改广告成功！", c)
}
