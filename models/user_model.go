package models

import (
	"GoBlog/models/ctype"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	NickName       string           `gorm:"size:36" json:"nickName"`
	Username       string           `gorm:"size:36" json:"username"`
	Password       string           `gorm:"size:128" json:"password"`
	Avatar         string           `gorm:"size:256" json:"avatar"`
	Email          string           `gorm:"size:128" json:"email"`
	Phone          string           `gorm:"size:18" json:"phone"`
	Addr           string           `gorm:"size:64" json:"addr"`
	Token          string           `gorm:"size:64" json:"token"`
	IP             string           `gorm:"size:20" json:"ip"`
	Role           ctype.Role       `gorm:"size:4,default:1" json:"role"`
	SignStatus     ctype.StatusType `gorm:"type=smallint(6)" json:"sign_Status"`
	ArticleModels  []ArticleModel   `gorm:"foreignKey:UserID" json:"-"`
	CollectsModels []ArticleModel   `gorm:"many2many:user_collect_models;joinForeignKey:UserId" json:"-"`
}
