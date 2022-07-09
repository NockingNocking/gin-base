/*
 * @Author: Mocking 497773732@qq.com
 * @Date: 2022-07-07 15:39:45
 * @LastEditors: Mocking 497773732@qq.com
 * @LastEditTime: 2022-07-07 16:35:53
 * @FilePath: \ginx\contraollers\userController.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package user

import (
	"ginx/contraollers/base"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	base.BaseController
}

func (con UserController) GetInfo(ctx *gin.Context) {
	// 获取中间件传过来的数据,获取到的数据是一个空接口类型
	changeName, _ := ctx.Get("sign")
	v, ok := changeName.(string) // 类型断言
	if ok {
		ctx.String(http.StatusOK, v)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"name": "张三",
			"age":  18,
			"sex":  "男",
		})
	}

	// 这里可以使用继承的base类的方法
	// con.Success(ctx)

}

// 增加用户信息
type USerInfos struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Sex      string `json:"sex" form:"sex"`
}

func (con UserController) AddUser(ctx *gin.Context) {
	var USerinfos USerInfos
	ctx.ShouldBind(&USerinfos)
	ctx.JSON(http.StatusOK, gin.H{
		"data": USerinfos,
	})
}
