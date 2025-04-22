package user_api

import (
	"GoBlog/models"
	"GoBlog/models/res"
	"GoBlog/service/common"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserList(c *gin.Context) {
	//参数校验
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}
	//分页逻辑
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
		Debug:    false,
	})
	res.OkWithList(list, count, c)
	return
}
