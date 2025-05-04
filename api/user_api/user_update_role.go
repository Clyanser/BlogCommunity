package user_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/ctype"
	"GoBlog/models/res"
	"github.com/gin-gonic/gin"
)

type UserRole struct {
	Role   ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	UserID uint       `json:"user_id" binding:"required" msg:"用户id错误"`
}

// UserUpdateRole 权限修改
func (UserApi) UserUpdateRole(c *gin.Context) {
	var cr UserRole
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var user models.UserModel
	err := global.DB.Take(&user, cr.UserID).Error
	if err != nil {
		res.FailWithMsg("用户id不存在", c)
		return
	}
	err = global.DB.Model(&user).Update("role", cr.Role).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("权限修改失败", c)
		return
	}
	res.OkWithMsg("修改权限成功", c)
}
