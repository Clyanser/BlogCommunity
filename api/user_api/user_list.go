package user_api

import (
	"GoBlog/models"
	"GoBlog/models/ctype"
	"GoBlog/models/res"
	"GoBlog/service/common"
	"GoBlog/utils/desens"
	"GoBlog/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserList(c *gin.Context) {
	//如何判断是否是管理员 - 根据token
	token := c.Request.Header.Get("Token")
	if token == "" {
		res.FailWithMsg("未携带token", c)
		return
	}
	//校验token的值
	claims, err := jwts.ParseToken(token)
	if err != nil {
		res.FailWithMsg("token错误", c)
		return
	}

	//参数校验
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}
	var users []models.UserModel
	//分页逻辑
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
		Debug:    false,
	})

	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			//	管理员
			user.Username = ""
		}
		//数据脱敏
		user.Phone = desens.DesensitizationTel(user.Phone)
		user.Email = desens.DesensitizationEmail(user.Email)
		users = append(users, user)
	}

	res.OkWithList(users, count, c)
	return
}
