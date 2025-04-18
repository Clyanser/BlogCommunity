package menu_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (MenuAPI) MenuUpdate(c *gin.Context) {
	var cr MenuCreateReq
	//参数校验
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//先把之前的banner 清空
	id := c.Param("id")
	var menuModel models.MenuModel
	err = global.DB.Preload("Banners").Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	//	清空逻辑
	global.DB.Model(&menuModel).Association("Banners").Clear()
	//	如果选择了banner 就添加
	if len(cr.ImageSortList) > 0 {
		//	操作第三张表
		var bannerList []models.MenuBannerModel
		for _, sort := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{
				MenuID:   menuModel.ID,
				BannerID: sort.ImageID,
				Sort:     sort.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
		}
		res.FailWithMsg("创建菜单图片失败", c)
		return
	}
	//	普通更新
	maps := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改菜单失败", c)
		return
	}
	res.FailWithMsg("修改菜单成功！", c)

}
