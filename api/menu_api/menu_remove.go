package menu_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (MenuAPI) MenuDelete(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	//使用事务逻辑
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		//删除逻辑
		err = tx.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error(err)
			return err
		}
		err = tx.Delete(&menuList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.OkWithMsg("删除菜单失败", c)
		return
	}
	//删除第三张表
	res.OkWithMsg(fmt.Sprintf("共删除 %d 个菜单", count), c)
}
