package user_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"GoBlog/utils/jwts"
	"GoBlog/utils/pwd"
	"github.com/gin-gonic/gin"
)

type UpdateUserPassword struct {
	OldPassword string `json:"old_password" binding:"required" msg:"请输入旧密码"` //旧密码
	NewPassword string `json:"new_password" binding:"required" msg:"请输入新密码"` //新密码
}

// UserUpdatePassword 修改登录人的id
func (UserApi) UserUpdatePassword(c *gin.Context) {
	_cliams, _ := c.Get("claims")
	claims := _cliams.(*jwts.CustomClaims)
	var cr UpdateUserPassword
	if err := c.ShouldBind(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	//判断密码是否一致
	if !pwd.VerifyPassword(user.Password, cr.OldPassword) {
		res.FailWithMsg("密码错误", c)
		return
	}
	hashPassword, _ := pwd.HashPwd(cr.NewPassword)
	err = global.DB.Model(&user).Update("password", hashPassword).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("密码修改失败！", c)
		return
	}
	res.OkWithMsg("密码修改成功", c)
	return
}
