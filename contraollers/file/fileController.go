/*
 * @Author: Mocking 497773732@qq.com
 * @Date: 2022-07-07 16:53:47
 * @LastEditors: Mocking 497773732@qq.com
 * @LastEditTime: 2022-07-08 15:30:59
 * @FilePath: \ginx\contraollers\file\fileController.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE


 */
package file

import (
	"fmt"
	"ginx/config"
	"ginx/contraollers/base"
	fileService "ginx/service/file"
	"ginx/utils"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	base.BaseController
}

// 单文件上传
func (con FileController) FileUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		utils.HandleError(err)
	}
	if !config.Cfg.OssIsOpen {
		result, ok := fileService.UploadFileToLocal(ctx, file)
		if ok {
			con.Success(result, ctx)
		}
		con.Fail(gin.Error{}, ctx)
	}

	result, ok := fileService.UploadFileToOss(ctx, file)
	if ok {
		con.Success(result, ctx)
	}
	con.Fail(gin.Error{}, ctx)
}

// 多文件上传
func (con FileController) FilesUpload(ctx *gin.Context) {

	form, err := ctx.MultipartForm()
	if err != nil {
		con.Fail(err, ctx)
		return
	}
	files := form.File["files"]

	for _, file := range files {
		fmt.Println(file.Filename)
		error := ctx.SaveUploadedFile(file, "./public/images/"+file.Filename)
		if error != nil {
			con.Fail(error, ctx)
		}
	}
	con.Success(len(files), ctx)
}

// 删除文件
func (con FileController) FileDelate(ctx *gin.Context) {
	fileName := ctx.Query("filename")
	res := fileService.DelateFileFromOss(fileName)
	con.Success(res, ctx)
}
