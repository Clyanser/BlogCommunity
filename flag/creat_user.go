package flag

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/ctype"
	"GoBlog/utils/pwd"
	"fmt"
)

func CreateUser(permissions string) {
	//	创建用户的逻辑
	//	用户名：昵称、密码、确认密码、邮箱
	var (
		userName        string
		nickName        string
		password        string
		confirmPassword string
		email           string
	)
	fmt.Printf("请输入用户名！：")
	fmt.Scanln(&userName)
	fmt.Printf("请输入昵称！：")
	fmt.Scanln(&nickName)
	fmt.Printf("请输入密码！：")
	fmt.Scanln(&password)
	fmt.Printf("请确认密码：")
	fmt.Scanln(&confirmPassword)
	fmt.Printf("请输入邮箱！：")
	fmt.Scanln(&email)

	//	判断用户名是否存在
	var usermodel models.UserModel
	err := global.DB.Take(&usermodel, "user_name = ?", userName).Error
	if err == nil {
		global.Log.Error("用户已存在,请重新输入！")
		return
	}
	//	校验两次密码
	if password != confirmPassword {
		global.Log.Error("两次输入的密码,不一致，请重新输入")
		return
	}
	//	对密码进行hash
	hash := pwd.HashPwd(password)
	//默认为普通用户
	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}

	//	头像问题
	//	1.默认头像2.随机选择头像
	avatar := "/uploads/avatar/default.png"
	//	入库
	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		Username:   userName,
		Password:   hash,
		Email:      email,
		Role:       role,
		Avatar:     avatar,
		IP:         "127.0.0.19",
		Addr:       "amoy",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户%s创建成功", userName)
}
