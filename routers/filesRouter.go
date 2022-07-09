package routers

import (
	"ginx/contraollers/file"

	"github.com/gin-gonic/gin"
)

func FileRouterInit(r *gin.Engine) {
	fileRouters := r.Group("/file")
	{
		fileRouters.POST("/upload", file.FileController{}.FileUpload)
		fileRouters.GET("/delate", file.FileController{}.FileDelate)
		fileRouters.POST("/uploads", file.FileController{}.FilesUpload)
	}
}
