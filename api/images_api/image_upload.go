package images_api

import (
	"GoBlog/global"
	"GoBlog/models"
	"GoBlog/models/ctype"
	"GoBlog/models/res"
	"GoBlog/plugins/qiniu"
	"GoBlog/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"path"
	"strings"
)

// 图片上传的白名单
var (
	WitheImagesList = []string{
		"jpg", "jpeg", "png", "gif",
	}
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
		//判断上传的图片后缀是否合法
		fileName := file.Filename
		nameList := strings.Split(fileName, ".")
		suffix := strings.ToLower(nameList[len(nameList)-1])
		if !utils.IsInList(suffix, WitheImagesList) {
			reslist = append(reslist, FileUploadResponse{
				Filename:  file.Filename,
				IsSuccess: false,
				Message:   "非法文件格式",
			})
			continue
		}

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

		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		byteData, err := io.ReadAll(fileObj)
		imageHash := utils.Md5V(byteData)

		//去数据库里面查，这个文件是否存在-根据hashValue
		var bannerModel models.BannerModel
		err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
		if err == nil {
			reslist = append(reslist, FileUploadResponse{
				Filename:  bannerModel.Path,
				IsSuccess: false,
				Message:   "图片重复！",
			})
			continue
		}
		//判断是否使用七牛云
		if global.Config.QiNiu.Enable {
			if filePath, err = qiniu.UploadImage(byteData, fileName, "goBlog"); err != nil {
				global.Log.Error(err)
				continue
			}
			reslist = append(reslist, FileUploadResponse{
				Filename:  filePath,
				IsSuccess: true,
				Message:   "七牛云上传成功！！",
			})
			global.DB.Create(&models.BannerModel{
				Path:      filePath,
				Hash:      imageHash,
				Name:      fileName,
				ImageType: ctype.QiNiu,
			})
			continue
		}

		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		reslist = append(reslist, FileUploadResponse{
			Filename:  filePath,
			IsSuccess: true,
			Message:   "上传成功！",
		})
		//图片入库
		global.DB.Create(&models.BannerModel{
			Path:      filePath,
			Hash:      imageHash,
			Name:      fileName,
			ImageType: ctype.Local,
		})
	}
	res.OkWithData(reslist, c)
	//fileheader, err := c.FormFile("image")
	//if err != nil {
	//	res.FailWithMsg(err.Error(), c)
	//	return
	//}

}
