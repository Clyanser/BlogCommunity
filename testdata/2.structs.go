package main

//
//import (
//	"fmt"
//	"github.com/fatih/structs"
//	"gorm.io/gorm"
//)
//
//type AdvertReq struct {
//	gorm.Model `structs:"-"`
//	Title      string `json:"title" binding:"required" msg:"请输入标题" structs:"title"`         //显示的标题
//	Href       string `json:"href" binding:"required,url" msg:"非法的跳转链接"structs:"-"`         //跳转链接
//	Images     string `json:"images" binding:"required,url" msg:"图片地址非法"structs:"-"`        //图片
//	IsShow     bool   `json:"is_show" binding:"required" msg:"请选择标题是否展示" structs:"is_show"` //是否展示
//}
//
//func main() {
//	u1 := AdvertReq{
//		Title:  "XXX",
//		Href:   "http://www.baidu.com",
//		Images: "XXX",
//		IsShow: true,
//	}
//	m3 := structs.Map(&u1)
//	fmt.Println(m3)
//}
