package images_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ImagesAPI) ImageRemove(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	var ImageList []models.BannerModel
	count := global.DB.Find(&ImageList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("文件不存在", c)
		return
	}
	//删除逻辑
	global.DB.Delete(&ImageList)
	res.OkWithMsg(fmt.Sprintf("共删除 %d 张图片", count), c)
}
