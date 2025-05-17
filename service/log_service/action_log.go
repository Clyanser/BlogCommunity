package log_service

import (
	"GoBlog/core"
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/ctype"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type ActionLog struct {
	c            *gin.Context
	level        ctype.LogLevelType
	title        string
	requestBody  []byte
	responseBody []byte
}

func (ac *ActionLog) SetTitle(title string) {
	ac.title = title
}
func (ac *ActionLog) SetLevel(level ctype.LogLevelType) {
	ac.level = level
}

func (ac *ActionLog) SetRequestBody(c *gin.Context) {
	//请求中间件
	byteData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		global.Log.Error(err.Error())
	}
	fmt.Printf("body: %s", string(byteData))
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byteData))
	ac.requestBody = byteData
}

func (ac *ActionLog) SetResponseBody(data []byte) {
	ac.responseBody = data
}

func (ac ActionLog) Save(userModel models.UserModel) {
	ip := ac.c.ClientIP()
	addr := core.GetIpAddr(ip)

	err := global.DB.Create(&models.LogModel{
		LogType: ctype.ActionType,
		Title:   ac.title,
		Content: "",
		Level:   ac.level,
		UserID:  &userModel.ID,
		Ip:      ip,
		Addr:    addr,
	}).Error
	if err != nil {
		global.Log.Errorf("日志创建失败 %s", err.Error())
		return
	}

}

func NewActionLog(c *gin.Context) *ActionLog {
	return &ActionLog{
		c: c,
	}
}

//func GetLog(c *gin.Context) *ActionLog {
//	_log, ok := c.Get("log")
//	if ok != nil {
//		return NewActionLog(c)
//	}
//	log, ok := _log.(*ActionLog)
//	if ok {
//		return NewActionLog(c)
//	}
//	return log
//}
