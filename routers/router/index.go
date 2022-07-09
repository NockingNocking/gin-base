/*
 * @Author: Mocking 497773732@qq.com
 * @Date: 2022-07-07 23:17:09
 * @LastEditors: Mocking 497773732@qq.com
 * @LastEditTime: 2022-07-08 11:50:01
 * @FilePath: \ginx\routers\index.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package router

import (
	"ginx/middlewares"
	"ginx/routers"
	"ginx/utils"
	"html/template"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func GinInit() {
	// 禁用控制台颜色
	//gin.DisableConsoleColor()
	// gin.ForceConsoleColor()

	//gin.New()返回一个*Engine 指针
	//而gin.Default()不但返回一个*Engine 指针，而且还进行了debugPrintWARNINGDefault()和engine.Use(Logger(), Recovery())其他的一些中间件操作
	Router = gin.Default()
	//Router = gin.New()
}

func SetupRouter(projectPath string) {

	// 加载自定义渲染模板函数
	Router.SetFuncMap(template.FuncMap{
		"UnixToTime": utils.UnixToTime,
	})
	// 加载全局模板
	Router.LoadHTMLGlob("templates/**/*")
	// 配置静态服务
	Router.Static("/public", "./public")
	Router.NoMethod(middlewares.InitNoMethodJson)
	Router.NoRoute(middlewares.InitNoRouteJson)
	Router.Use(middlewares.InitErrorHandler)
	Router.Use(middlewares.InitAccessLogMiddleware)

	// 路由分组(用户板块路由)
	routers.UserRouterInit(Router)
	// 文件上传
	routers.FileRouterInit(Router)
}
