package main

import (
	"todo-list/config"
	"todo-list/model"
	"todo-list/router"
)

func main() {
	// 1. 初始化数据库连接
	config.InitDB()
	// 2. 自动迁移表结构（创建todo表）
	model.AutoMigrate(config.DB)
	// 3. 初始化路由
	r := router.InitRouter()
	// 4. 启动服务（默认8080端口）
	err := r.Run(":8080")
	if err != nil {
		panic("服务启动失败：" + err.Error())
	}
}
