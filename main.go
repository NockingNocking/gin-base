/*
 * @Author: Mocking 497773732@qq.com
 * @Date: 2022-07-07 11:00:12
 * @LastEditors: Mocking 497773732@qq.com
 * @LastEditTime: 2022-07-09 12:46:12
 * @FilePath: \ginx\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"flag"
	"ginx/config"
	"ginx/config/database"
	"ginx/logger"
	"ginx/routers/router"
)

var configFile = flag.String("configFile", "config/config.yml", "配置文件路径")
var projectPath = flag.String("projectPath", "/ginx", "项目访问路径前缀")

func init() {
	config.ConfigRead(*configFile)
	logger.LogInit()
	database.DatabaseInit()
	//TODO: 操作数据库 2022.07.09
}

func main() {

	router.GinInit()
	//gin工程实例 *gin.Engine
	r := router.Router
	//路由初始化
	router.SetupRouter(*projectPath)
	//监听端口
	r.Run(":" + config.Cfg.ListenPort)

}
