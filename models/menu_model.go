package models

import (
	"GoBlog/models/ctype"
	"gorm.io/gorm"
)

// MenuModel 菜单表 菜单的路径可以是 /path 也可以是路由别名
type MenuModel struct {
	gorm.Model
	Title        string        `gorm:"size:32" json:"title"`                                                                      // 菜单标题
	Path         string        `gorm:"size:32" json:"path"`                                                                       // 英文菜单标题
	Slogan       string        `gorm:"size:64" json:"slogan"`                                                                     // slogan
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract"`                                                               // 简介
	AbstractTime int           `json:"abstract_time"`                                                                             // 简介的切换时间
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;joinForeignKey:MenuID;JoinReferences:BannerID" json:"banners"` // 菜单的图片列表
	BannerTime   int           `json:"banner_time"`                                                                               // 菜单图片的切换时间 为 0 表示不切换
	Sort         int           `gorm:"size:10" json:"sort"`                                                                       // 菜单的顺序
}
