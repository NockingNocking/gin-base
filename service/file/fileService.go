/*
 * @Author: Mocking 497773732@qq.com
 * @Date: 2022-07-08 14:14:16
 * @LastEditors: Mocking 497773732@qq.com
 * @LastEditTime: 2022-07-09 12:28:10
 * @FilePath: \ginx\service\fileService.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package fileService

import (
	"bytes"
	"ginx/config"
	"ginx/config/alioss"
	"ginx/utils"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 上传到本地的逻辑
func UploadFileToLocal(ctx *gin.Context, file *multipart.FileHeader) (string, bool) {
	extName := path.Ext(file.Filename)
	if res, ok := utils.CheckFileExt(extName); !ok {
		return res, ok
	}

	dir := "./public/images/" + utils.GetDay()
	if err := os.MkdirAll(dir, 0666); err != nil {
		utils.HandleError(err)
	}
	fileName := strconv.FormatInt(utils.GetUnix(), 10) + extName
	dst := path.Join(dir, fileName)
	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		utils.HandleError(err)
	}

	return dst, true
}

// TODO:这里可以把 bucketName 不用写固定
var bucketName string = "nocking-blog-artical"

// 上传到oss的逻辑
func UploadFileToOss(ctx *gin.Context, file *multipart.FileHeader) (string, bool) {
	extName := path.Ext(file.Filename)
	if res, ok := utils.CheckFileExt(extName); !ok {
		return res, ok
	}

	fileHandle, err := file.Open() //打开上传文件
	if err != nil {
		utils.HandleError(err)
	}
	defer fileHandle.Close()

	fileByte, err := ioutil.ReadAll(fileHandle) //获取上传文件字节流
	if err != nil {
		utils.HandleError(err)
	}

	// dst, _ := alioss.UploadToOss(file.Filename, "nocking-blog-artical", fileByte)

	domain := config.Cfg.OssDomain

	bucket, err := alioss.GetBucket(bucketName)
	if err != nil {
		utils.HandleError(err)
	}

	folderName := utils.GetDay()
	yunFileTmpPath := path.Join("blog", folderName+"/"+file.Filename)

	err = bucket.PutObject(yunFileTmpPath, bytes.NewReader([]byte(fileByte)))
	if err != nil {
		utils.HandleError(err)
	}

	return domain + yunFileTmpPath, true
}

// 从oss删除图片
// fileName string
func DelateFileFromOss(fileName string) string {
	bucket, err := alioss.GetBucket(bucketName)
	if err != nil {
		utils.HandleError(err)
	}
	err = bucket.DeleteObject(fileName)
	if err != nil {
		utils.HandleError(err)
	}

	return "删除成功！"
}
