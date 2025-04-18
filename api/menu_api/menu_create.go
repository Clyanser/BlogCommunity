package menu_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/ctype"
	"GoBlog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuCreateReq struct {
	Title         string      `json:"title" binding:"required" msg:"请输入菜单名称"`
	Path          string      `json:"path" binding:"required" msg:"请输入菜单路径"`
	Slogan        string      `json:"slogan"`
	Abstract      ctype.Array `json:"abstract"`
	AbstractTime  int         `json:"abstract_time"`                                //切换的时间，单位秒
	BannerTime    int         `json:"banner_time"`                                  //切换的时间，单位秒
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号"` //排序
	ImageSortList []ImageSort `json:"image_sort_list"`                              //具体图片的顺序
}

func (MenuAPI) MenuCreate(c *gin.Context) {
	var cr MenuCreateReq
	//参数校验
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//	重复值的判断
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "title = ? or path = ?", cr.Title, cr.Path).RowsAffected
	fmt.Println(count)
	if count > 0 {
		res.FailWithMsg("存在重复的菜单", c)
		return
	}
	menumodels := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	//	添加逻辑
	//	1。创建banner数据入库
	err = global.DB.Create(&menumodels).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单添加失败", c)
		return
	}
	//	批量入库
	var menuBannerList []models.MenuBannerModel
	if len(cr.ImageSortList) == 0 {
		res.OkWithMsg("菜单添加成功！", c)
		return
	}
	for _, sort := range cr.ImageSortList {
		//这里需要判断Image_ID图片的是否真的存在
		menuBannerList = append(menuBannerList, models.MenuBannerModel{

			MenuID:   menumodels.ID,
			BannerID: uint(sort.ImageID),
			Sort:     sort.Sort,
		})
	}
	//2.给第三张表入库
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("图片关联失败", c)
		return
	}
	res.OkWithMsg("菜单添加成功！", c)
}
