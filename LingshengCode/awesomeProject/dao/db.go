package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var Db *gorm.DB

func init() {
	// 配置 MySQL
	username := "root"
	password := "123456"
	host := "pogf.com.cn"
	port := 3306
	dbname := "tiktok"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("数据库连接失败", err)
	}
	db.AutoMigrate(&UserBasicInfo{})
	Db = db
}
