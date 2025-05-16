package log_service

import (
	"GoBlog/core"
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/ctype"
	"GoBlog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewLoginSuccess(c *gin.Context, loginType ctype.StatusType, userModel models.UserModel) {
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)

	token := c.GetHeader("token")
	fmt.Println(token)

	err := global.DB.Create(&models.LogModel{
		LogType:     ctype.LoginLogType,
		Title:       "用户登录",
		Content:     "登录成功",
		UserID:      &userModel.ID,
		Ip:          ip,
		Addr:        addr,
		LoginStatus: true,
		Username:    userModel.Username,
		Password:    "-",
		LoginType:   loginType,
	}).Error
	if err != nil {
		global.Log.Errorf("写入日志失败 %s", err.Error())
		res.FailWithMsg("写入日志失败", c)
		return
	}
}

func NewLoginFail(c *gin.Context, loginType ctype.StatusType, msg, username, pwd string) {
	ip := c.ClientIP()
	addr := core.GetIpAddr(ip)

	err := global.DB.Create(&models.LogModel{
		LogType:     ctype.LoginLogType,
		Title:       "用户登录失败",
		Content:     msg,
		Ip:          ip,
		Addr:        addr,
		LoginStatus: false,
		Username:    username,
		Password:    pwd,
		LoginType:   loginType,
	}).Error
	if err != nil {
		global.Log.Errorf("写入日志失败 %s", err.Error())
		res.FailWithMsg("写入日志失败", c)
		return
	}
}
