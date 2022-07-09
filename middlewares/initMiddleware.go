/*
 * @Author: Mocking 497773732@qq.com
 * @Date: 2022-07-07 16:22:11
 * @LastEditors: Mocking 497773732@qq.com
 * @LastEditTime: 2022-07-08 11:46:16
 * @FilePath: \ginx\middlewares\testMiddleware.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package middlewares

import (
	"encoding/json"
	"fmt"
	"ginx/logger"
	"ginx/threadlocal"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/jtolds/gls"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*
	中间件中使用goroutine  不能直接使用原始的上下文(ctx *gin.Context),
	需要使用其只读副本 (ctx.Copy())

*/

// 中间件和其他中间件以及控制器共享数据用ctx.Set()  和 ctx.Get()
func InitMiddleware(ctx *gin.Context) {
	ctx.Set("sign", "王二麻子")
	// 请求时打印当前时间
	fmt.Println(time.Now())
}

// 返回请求错误
func InitErrorHandler(ctx *gin.Context) {
	ctx.Next()
	if len(ctx.Errors) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": ctx.Errors,
		})
	}
}

//未知路由处理 返回json
func InitNoRouteJson(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"code": http.StatusNotFound,
		"msg":  "path not found",
	})

}

//未知调用方式 返回json
func InitNoMethodJson(ctx *gin.Context) {
	ctx.JSON(http.StatusMethodNotAllowed, gin.H{
		"code": http.StatusMethodNotAllowed,
		"msg":  "method not allowed",
	})
}

//打印请求和响应日志
func InitAccessLogMiddleware(ctx *gin.Context) {
	//request id
	requestId := ctx.Request.Header.Get("X-RequestId")
	if requestId == "" {
		requestId = strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	//response requestId
	ctx.Writer.Header().Set("X-RequestId", requestId)

	// 开始时间
	startTime := time.Now()

	//处理请求 do chian
	threadlocal.Mgr.SetValues(gls.Values{threadlocal.Rid: requestId}, func() {
		ctx.Next()
	})

	// 结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)
	// 请求方式
	reqMethod := ctx.Request.Method
	// 请求路由
	reqUri := ctx.Request.RequestURI
	// 状态码
	statusCode := ctx.Writer.Status()
	// 请求IP
	clientIP := ctx.ClientIP()
	//请求参数
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	//返回参数
	responseMap := ctx.Keys
	responseJson, _ := json.Marshal(responseMap)

	// 日志格式
	logger.LogAccess.WithFields(logrus.Fields{

		"status_code":  statusCode,
		"latency_time": latencyTime,
		"client_ip":    clientIP,
		"req_method":   reqMethod,
		"req_uri":      reqUri,
		"req_Id":       requestId,
		"req_body":     string(body),
		"res_body":     string(responseJson),
	}).Info()
}
