package menu_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"github.com/gin-gonic/gin"
)

func (MenuAPI) MenuDetails(c *gin.Context) {
	//	先查菜单
	id := c.Param("id")
	var menuModel models.MenuModel
	err := global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	//查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id = ?", id)
	//解决null值问题
	banners := []Banner{}
	for _, banner := range menuBanners {
		if menuModel.ID != banner.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   banner.MenuID,
			Path: banner.BannerModel.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	res.OkWithData(menuResponse, c)
}
