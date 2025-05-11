package database

import (
	"Boke/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMySQL() {
	dsn := "root:123456@tcp(43.133.35.111:3306)/Article?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	err = DB.AutoMigrate(&model.Article{})
	if err != nil {
		panic("自动迁移失败: " + err.Error())
	}

	fmt.Println("数据库初始化成功")
}
