package user_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"GoBlog/utils/jwts"
	"GoBlog/utils/pwd"
	"github.com/gin-gonic/gin"
)

type EmailLoginRequuest struct {
	Username string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (UserApi) EmailLogin(c *gin.Context) {
	//	参数绑定
	var cr EmailLoginRequuest
	//参数校验
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//验证用户名是否存在
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "username=? or email= ?", cr.Username, cr.Username).Error
	if err != nil {
		global.Log.Warn("用户名不存在")
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	//	校验密码
	isCheck := pwd.VerifyPassword(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	//	登录成功
	token, err := jwts.GetToken(jwts.JwtPayload{
		Nickname: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("token生成失败", c)
		return
	}
	res.OkWithData(token, c)
}
