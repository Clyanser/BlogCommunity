package menu_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	MenuModel models.MenuModel
	Banners   []Banner
}

func (MenuAPI) MenuList(c *gin.Context) {
	//	先查菜单
	var menuList []models.MenuModel
	var menuIDList []uint
	global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)
	fmt.Println(menuList, menuIDList)
	//查连接表
	var menuBanners []models.MenuBannerModel
	global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBanners, "menu_id in ?", menuIDList)
	var menuResponse []MenuResponse
	for _, model := range menuList {
		//model就是一个菜单
		//解决null值问题
		banners := []Banner{}
		for _, banner := range menuBanners {
			if model.ID != banner.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   banner.MenuID,
				Path: banner.BannerModel.Path,
			})
		}
		menuResponse = append(menuResponse, MenuResponse{
			MenuModel: model,
			Banners:   banners,
		})
	}
	res.OkWithData(menuResponse, c)
}
