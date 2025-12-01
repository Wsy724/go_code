package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局DB对象
var DB *gorm.DB

// 初始化数据库连接
func InitDB() {
	// dsn格式：用户名:密码@tcp(地址:端口)/数据库名?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:Wsy1817396846@tcp(127.0.0.1:3306)/todo_list?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败：" + err.Error())
	}
}
