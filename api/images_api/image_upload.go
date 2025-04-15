package images_api

import (
	"GoBlog/global"
	"GoBlog/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
)

// 构造一个图片上传的响应
type FileUploadResponse struct {
	Filename  string `json:"file_name"`  //文件名
	Message   string `json:"message"`    //消息
	IsSuccess bool   `json:"is_success"` //是否上传成功
}

// 上传单个图片，返回图片的url
func (ImagesAPI) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMsg("不存在的错误文件", c)
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMsg("不存在的错误文件", c)
		return
	}
	var reslist []FileUploadResponse

	//判断路径是否存在，不存在就创建

	for _, file := range fileList {
		filePath := path.Join(global.Config.Upload.Path, file.Filename)
		//判断大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			reslist = append(reslist, FileUploadResponse{
				Filename:  file.Filename,
				IsSuccess: false,
				Message:   fmt.Sprintf("图片大小超过上限,当前大小为:%dMB 上限为：%dMB", size, global.Config.Upload.Size),
			})
			continue
		}
		reslist = append(reslist, FileUploadResponse{
			Filename:  filePath,
			IsSuccess: true,
			Message:   "上传成功！",
		})
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			continue
		}
	}
	res.OkWithData(reslist, c)
	//fileheader, err := c.FormFile("image")
	//if err != nil {
	//	res.FailWithMsg(err.Error(), c)
	//	return
	//}

}
