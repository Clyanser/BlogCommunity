package tag_api

import (
	"GoBlog/models"
	"GoBlog/models/res"
	"GoBlog/service/common"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagList(c *gin.Context) {
	var cr models.PageInfo
	if err := c.ShouldBind(&cr); err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}
	//	分页查询
	list, count, _ := common.ComList(models.TagModel{}, common.Option{
		PageInfo: cr,
	})
	//需要展示这个标签下文章的数量
	res.OkWithList(list, count, c)

}
