package common

import (
	"GoBlog/global"
	"GoBlog/models"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MySqlLog})
	}
	if option.Page <= 0 {
		option.Page = 1
	}
	if option.Limit <= 0 {
		option.Limit = 10
	}
	//排序默认值
	if option.Sort == "" {
		option.Sort = "created_at desc" //默认按照时间往前排
	}
	//查所有的图片列表
	count = global.DB.Select("id").Find(&list).RowsAffected
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	//列表查询分页
	err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return list, count, err
}
