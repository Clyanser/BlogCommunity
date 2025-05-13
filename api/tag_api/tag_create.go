package tag_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"github.com/gin-gonic/gin"
)

type TagReq struct {
	Title string `json:"title" binding:"required" msg:"请输入标题" structs:"title"` //显示的标题
}

func (TagApi) TagCreate(c *gin.Context) {
	var cr TagReq
	//参数校验
	err := c.ShouldBind(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//重复的判断
	var tag models.TagModel
	err = global.DB.Take(&tag, "title = ?", cr.Title).Error
	if err == nil {
		res.FailWithMsg("该标签已存在", c)
		return
	}
	//添加逻辑
	err = global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("添加标签失败！", c)
		return
	}
	res.OkWithMsg("添加标签成功！", c)
}
