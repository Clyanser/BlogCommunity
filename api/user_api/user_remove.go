package user_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (UserApi) UserRemove(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ParamsError, c)
		return
	}

	var userList []models.UserModel
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("用户不存在", c)
		return
	}
	//使用事务逻辑
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		//TODO:删除用户，还没实现消息表，评论表，用户收藏的文章，用户发布的文章 联合删除
		err = tx.Delete(&userList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.OkWithMsg("删除用户失败", c)
		return
	}
	//删除第三张表
	res.OkWithMsg(fmt.Sprintf("共删除 %d 个用户", count), c)
}
