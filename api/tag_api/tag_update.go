package tag_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagUpdate(c *gin.Context) {
	id := c.Param("id")
	var cr TagReq
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		res.FailWithMsg("标签不存在！", c)
		return
	}
	//结构体转map的第三方包
	maps := structs.Map(&cr)
	err = global.DB.Model(&tag).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改标签失败", c)
		return
	}
	res.FailWithMsg("修改标签成功！", c)
}
