package flag

import (
	"GoBlog/global"
	"GoBlog/models/ctype"
	"GoBlog/service/user_ser"
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
	//	校验两次密码
	if password != confirmPassword {
		global.Log.Error("两次输入的密码,不一致，请重新输入")
		return
	} else {
		global.Log.Infof("输入的密码: [%s]", password)
	}
	//默认为普通用户
	role := ctype.PermissionUser
	if permissions == "admin" {
		role = ctype.PermissionAdmin
	}
	err := user_ser.UserService{}.UserCreate(userName, nickName, permissions, role, email, password)
	if err != nil {
		global.Log.Error(err.Error())
		return
	}
	global.Log.Infof("用户%s创建成功", userName)
}
