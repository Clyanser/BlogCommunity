package user_api

import (
	"GoBlog/global"
	"GoBlog/models/ctype"
	"GoBlog/models/res"
	"GoBlog/service/user_ser"
	"github.com/gin-gonic/gin"
)

type userCreateRequest struct {
	NickName string     ` json:"nickName" binding:"required" msg:"请输入昵称"`
	Username string     ` json:"username" binding:"required" msg:"请输入用户名"`
	Password string     ` json:"password" binding:"required" msg:"请输入密码"`
	Role     ctype.Role ` json:"role" binding:"required" msg:"请选择角色"`
}

func (UserApi) UserCreate(c *gin.Context) {
	var cr userCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}
	err = user_ser.UserService{}.UserCreate(cr.Username, cr.NickName, cr.Password, cr.Role, "", c.ClientIP())
	if err != nil {
		global.Log.Error(err.Error())
		return
	}
	global.Log.Infof("用户%s创建成功", cr.Username)
	res.OkWithMsg("用户创建成功", c)
}
