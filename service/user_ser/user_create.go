package user_ser

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/ctype"
	"GoBlog/utils/pwd"
)

func (UserService) UserCreate(userName string, nickName, password string, role ctype.Role, email string, ip string) error {
	//	判断用户名是否存在
	var usermodel models.UserModel
	err := global.DB.Take(&usermodel, "username = ?", userName).Error
	if err == nil {
		global.Log.Error("用户已存在,请重新输入！")
		return err
	}
	//	对密码进行hash
	hash, _ := pwd.HashPwd(password)

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
		IP:         ip,
		Addr:       "amoy",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return err
	}
	global.Log.Infof("用户%s创建成功", userName)
	return nil
}
