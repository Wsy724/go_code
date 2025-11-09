package global

import (
	"fmt"
	"gorm_demo02/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:Wsy1817396846@tcp(127.0.0.1:3306)/gorm_db?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//DisableForeignKeyConstraintWhenMigrating: true, //禁止生成实体外键
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
}

func Migrate() {
	err := DB.AutoMigrate(&models.UserModel{}, &models.UserDetailModel{}, &models.GirlModel{}, &models.BoyModel{})
	if err != nil {
		log.Fatalf("数据库迁移失败%s", err)
		return
	} else {
		fmt.Println("数据库迁移成功~")
	}
}
