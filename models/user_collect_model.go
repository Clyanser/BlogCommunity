package models

import "time"

// 自定义第三张表，记录用户收藏的文章
type UserCollectModel struct {
	UserID       uint         `gorm:"foreignKey"`
	UserModel    UserModel    `gorm:"foreignKey:UserID"`
	ArticleID    uint         `gorm:"foreignKey"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
	CreatedAt    time.Time
}
