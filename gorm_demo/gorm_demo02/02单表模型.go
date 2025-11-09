package main

import (
	"fmt"
	"gorm_demo02/global"
	"gorm_demo02/models"
)

// 单表模型
func migrate() {
	err := global.DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		fmt.Println("表结构迁移失败~", err)
		return
	} else {
		fmt.Println("表结构迁移成功~")
	}
}

func main() {
	global.Connect()
	migrate()
}
