package user_api

import (
	"GoBlog/global"
	"GoBlog/models/res"
	"GoBlog/service"
	"GoBlog/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (UserApi *UserApi) Logout(c *gin.Context) {
	_cliams, _ := c.Get("claims")
	claims := _cliams.(*jwts.CustomClaims)

	token := c.Request.Header.Get("token")
	//需要计算距离现在的过期时间
	err := service.ServiceApp.UserService.Logout(claims, token)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("注销失败", c)
		return
	}
	res.OkWithMsg("注销成功！", c)
}
