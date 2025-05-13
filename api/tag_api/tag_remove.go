package tag_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagRemove(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	var TagList []models.TagModel
	count := global.DB.Find(&TagList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("标签不存在", c)
		return
	}
	//如果这个标签下有关联的文章怎么办？
	//删除逻辑
	global.DB.Delete(&TagList)
	res.OkWithMsg(fmt.Sprintf("共删除 %d 个标签", count), c)
}
