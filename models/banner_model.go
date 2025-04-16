package models

import (
	"GoBlog/global"
	"GoBlog/models/ctype"
	"gorm.io/gorm"
	"os"
)

type BannerModel struct {
	gorm.Model
	Path      string          `json:"path"`                        //图片路径
	Hash      string          `json:"hash"`                        //图片的hash值，用户判断是否重复图片
	Name      string          `gorm:"size:38" json:"name"`         //图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` //图片的上传类型，本地或者云
}

// 创建hook函数
func (b *BannerModel) BeforeDelete(tx *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		//	本地图片处理逻辑
		err = os.Remove(b.Path)
		if err != nil {
			global.Log.Error(err.Error())
			return err
		}
	}
	return nil
}
