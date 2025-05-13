package user_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"GoBlog/plugins/email"
	"GoBlog/utils/jwts"
	"GoBlog/utils/pwd"
	"GoBlog/utils/random"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password" `
}

func (UserApi) UserBindEmail(c *gin.Context) {
	_cliams, _ := c.Get("claims")
	claims := _cliams.(*jwts.CustomClaims)
	//用户绑定邮箱
	//第一次输入邮箱，后台会给这个邮箱发验证码，第二次用户输入邮箱，验证码，密码 完成绑定
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//判断验证码是否为空
	session := sessions.Default(c)
	if cr.Code == nil {
		// 第一次发验证码
		randomCode := random.RandomCode(4)
		session.Set("valid_code", randomCode)
		err := session.Save()
		if err != nil {
			res.FailWithMsg("session 保存失败: "+err.Error(), c)
			return
		}
		fmt.Println("验证码发送：", randomCode)
		email.NewCode().Send(cr.Email, "你的验证码为："+randomCode)
		res.OkWithMsg("验证码发送成功，请查收邮箱", c)
		return
	}

	// 第二次校验验证码，使用 *cr.Code 前要确保不是 nil
	code := session.Get("valid_code")
	if code == nil {
		res.FailWithMsg("验证码已过期或未发送", c)
		return
	}
	fmt.Printf("session code: %v, user input: %v\n", code, *cr.Code)

	if code != *cr.Code {
		res.FailWithMsg("验证码不正确", c)
		return
	}
	//修改用户邮箱
	var users models.UserModel
	err = global.DB.Take(&users, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	if len(cr.Password) < 4 {
		res.FailWithMsg("密码非法", c)
		return
	}
	//第一次的邮箱与第二次的邮箱也需要做一致性校验
	hashPassword := pwd.HashPwd(cr.Password)
	err = global.DB.Model(&users).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPassword,
	}).Error
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMsg("绑定邮箱失败", c)
	}
	res.OkWithMsg("邮箱绑定成功", c)
}
