package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id    int
	Name  string
	Hobby string
}

func main() {

	dsn := "root:Wsy1817396846@tcp(127.0.0.1:3306)/gorm_db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	var userList []User
	db.Find(&userList)
	for _, user := range userList {
		fmt.Println(user)
	}

}
