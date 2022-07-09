package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 公共处理方法类
type BaseController struct{}
type Data interface{}

func (con BaseController) Success(data Data, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

func (con BaseController) Fail(err error, ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code":    201,
		"message": err.Error(),
	})
}
