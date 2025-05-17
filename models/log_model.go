package models

import (
	"GoBlog/models/ctype"
	"gorm.io/gorm"
)

type LogModel struct {
	gorm.Model
	LogType     ctype.LogType      `json:"log_type"` //1-登录日志  2-操作日志  3-运行日志
	Title       string             `json:"title"`
	Content     string             `json:"content"`
	Level       ctype.LogLevelType `json:"level"`
	UserID      *uint              `json:"user_id"`
	UserModel   UserModel          `gorm:"foreignKey:ID" json:"-"`
	Ip          string             `json:"ip"`
	Addr        string             `json:"addr"`
	LoginStatus bool               `json:"login_status"`
	Username    string             `gorm:"size:36" json:"username"`
	Password    string             `gorm:"size:128" json:"password"`
	LoginType   ctype.StatusType   `json:"login_type"`
}
