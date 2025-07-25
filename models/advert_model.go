package models

import "gorm.io/gorm"

type AdvertModel struct {
	gorm.Model
	Title  string `gorm:"size:32" json:"title"` //显示的标题
	Href   string `json:"href"`                 //跳转链接
	Images string `json:"images"`               //图片
	IsShow bool   `json:"is_show"`              //是否展示
}
