/*
 * @Author: Mocking 497773732@qq.com
 * @Date: 2022-07-08 21:47:09
 * @LastEditors: Mocking 497773732@qq.com
 * @LastEditTime: 2022-07-09 12:23:03
 * @FilePath: \ginx\models\sql\sqlInit.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package database

import (
	"fmt"
	"ginx/utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseInit() *gorm.DB {

	dsn := "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"

	// dsn := config.Cfg.MysqlUser + ":" + config.Cfg.MysqlPwd + "@tcp" + "(" + config.Cfg.MysqlHost + ":" + config.Cfg.MysqlPort + ")"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.HandleError(err)
	}
	fmt.Println(err, "--------")
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
