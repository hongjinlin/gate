/*
@Time : 2022/10/4 16:14
@Author : hongjinlin
@File : db
@Software: GoLand
*/

package sys

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func init() {
	dsn := DSN + "?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("db error %s \n", err.Error())
	}
}
