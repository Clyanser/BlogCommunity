package advert_api

import (
	"GoBlog/models"
	"GoBlog/models/res"
	"GoBlog/service/common"
	"github.com/gin-gonic/gin"
	"strings"
)

// AdvertList 广告列表
// @Summary 广告列表
// @Description 广告列表
// @Tags 广告管理
// @Param data body models.PageInfo  true "表示多个参数"
// @Router /api/adverts [get]
// produce json
// @Success 200 {object} res.Response{}
func (AdvertApi) AdvertList(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBind(&cr); err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}
	//判断Referer,是否包含admin,如果是就全部返回，不是就返回is_show=true
	isShow := true
	referer := c.GetHeader("Referer")
	if strings.Contains(referer, "admin") {
		isShow = false
	}
	//	分页查询
	list, count, _ := common.ComList(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	res.OkWithList(list, count, c)

}
