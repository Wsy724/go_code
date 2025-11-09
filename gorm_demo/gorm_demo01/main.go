package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// UserInfo --> UserInfo数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// 替换为你的实际数据库信息
	//db, err := gorm.Open("mysql", "root:Wsy1817396846@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//
	////创建表 自动迁移 (把结构体和数据表进行对应)
	//db.AutoMigrate(&UserInfo{})
	//
	////创建数据行
	//u1 := UserInfo{1, "唐僧", "男", "念经"}
	//db.Create(&u1)

	dsn := "root:Wsy1817396846@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(UserInfo{})

	mysqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("获取数据库实例失败: %v\n", err)
	}
	//
	defer mysqlDB.Close()

	//u2 := UserInfo{3, "八戒", "男", "找嫦娥"}
	//db.Create(&u2)

	//查询
	var u UserInfo
	db.First(&u, 1)
	fmt.Printf("%#v\n", u)

	//更新
	db.Model(&u).Update("hobby", "取经")

}
