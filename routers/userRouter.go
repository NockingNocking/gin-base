/*
 * @Author: Mocking 497773732@qq.com
 * @Date: 2022-07-07 12:15:21
 * @LastEditors: Mocking 497773732@qq.com
 * @LastEditTime: 2022-07-07 23:14:13
 * @FilePath: \ginx\routers\userRouter.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package routers

import (
	"ginx/contraollers/user"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(r *gin.Engine) {
	userRouters := r.Group("/user")
	// 路由分组中间件
	// userRouters.Use(middlewares.InitMiddleware)
	{
		userRouters.GET("/info", user.UserController{}.GetInfo)
		userRouters.POST("/add", user.UserController{}.AddUser)
	}

}
