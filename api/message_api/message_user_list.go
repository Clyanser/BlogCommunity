package message_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/utils/jwts"
	"github.com/gin-gonic/gin"
)

func (MessageAPI) MessageUserList(c *gin.Context) {
	_cliams, _ := c.Get("claims")
	claims := _cliams.(*jwts.CustomClaims)
	var messageList []models.MessageModel
	global.DB.Find(&messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)

}
