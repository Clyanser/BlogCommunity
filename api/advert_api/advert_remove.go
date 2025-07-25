package advert_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

// AdvertRemove 删除广告
// @Summary 删除广告
// @Description 删除广告
// @Tags 广告管理
// @Param data query models.RemoveRequest true "广告id列表"
// @Router /api/adverts [delete]
// @produce json
// @Accept json
// @Success 200 {object} res.Response{data=string}
func (AdvertApi) AdvertRemove(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	var AdvertList []models.AdvertModel
	count := global.DB.Find(&AdvertList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("广告不存在", c)
		return
	}
	//删除逻辑
	global.DB.Delete(&AdvertList)

	res.OkWithMsg(fmt.Sprintf("共删除 %d 个广告", count), c)
}
