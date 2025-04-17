package advert_api

import (
	"GoBlog/models"
	"GoBlog/models/res"
	"GoBlog/service/common"
	"github.com/gin-gonic/gin"
	"strings"
)

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
