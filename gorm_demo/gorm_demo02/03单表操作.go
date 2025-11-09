package main

import (
	"fmt"
	"gorm_demo02/global"
	"gorm_demo02/models"
	"time"
)

func insert() {
	//user01 := models.UserModel{
	//	Name: "悟空",
	//	Age:  18,
	//}

	userList := []models.UserModel{
		{Name: "八戒", Age: 22},
		{Name: "沙僧", Age: 23},
		{Name: "白龙马", Age: 20},
	}

	err := global.DB.Create(&userList).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(userList)
}

func query() {
	//查全部
	var userList1 []models.UserModel
	global.DB.Find(&userList1, "age>=?", 20) //条件查询
	fmt.Println(userList1)

	//查一个
	var user01 models.UserModel
	//根据主键去查询
	//global.DB.Take(&user01, 1)
	//fmt.Println(user01)
	//根据主键排序去查询第一个
	global.DB.Debug().First(&user01, 7)
	//global.DB.Debug().Last((&user01, 7)
	fmt.Println(user01)
}

func update() {
	var user = models.UserModel{Id: 3}
	global.DB.Model(&user).Update("age", 33)
	//global.DB.Model(&user).UpdateColumn("age", 33)//这个UpdateColumn不会走钩子函数
	fmt.Println(user)
}

func updates() {
	var user = models.UserModel{Id: 3}
	global.DB.Model(&user).Updates(models.UserModel{
		Name: "龙王",
		Age:  99,
	})
	fmt.Println(user)
}

func updateSave() {
	var user models.UserModel
	user.Id = 3
	user.Name = "八戒"
	user.Age = 22
	user.CreatedAt = time.Now()
	global.DB.Save(&user)
	fmt.Println(user)
}

func Delete() {
	var user = models.UserModel{Id: 3}
	global.DB.Delete(&user)
	//global.DB.Delete(&models.UserModel{Id: 3})
	//global.DB.Delete(&models.UserModel{}, "name=?","八戒")

	//批量删除
	//global.DB.Delete(&models.UserModel{},[]int{1,2,3,4})

	//跳过软删除DeletedAt，使用Unscoped()真正删除
	//global.DB.Unscoped().Delete(&user)

}

func Highquery() {
	//1.
	//var user []models.UserModel
	////条件查询
	//global.DB.Debug().Where("age>=?", 800).Find(&user)
	//for _, v := range user {
	//	fmt.Println(v.Name, v.Age, "岁")
	//}

	//2.where里面放结构体
	//var user models.UserModel
	//global.DB.Debug().Where(models.UserModel{
	//	Name: "唐僧",
	//	Age:  25,
	//}).Find(&user)
	//fmt.Println(user)

	//3.where里面放map
	//var user models.UserModel
	//global.DB.Debug().Where(map[string]any{
	//	"age": 800,
	//}).Find(&user)
	//fmt.Println(user)

	//4. 嵌套where条件
	//query01 := global.DB.Where("age > ?", 100) //条件1
	//var user []models.UserModel
	//global.DB.Debug().Where(query01).Find(&user)
	//fmt.Println(user)

	//or (满足任意一个都查出来)
	//var user []models.UserModel
	//global.DB.Debug().Or("age>=?", 800).Or("name=?", "唐僧").Find(&user)
	//fmt.Println(user)

	//not (排除条件的都查出来)
	//var user []models.UserModel
	//global.DB.Debug().Not("age>=?", 800).Find(&user)
	//fmt.Println(user)

	//order (排序)
	//var user []models.UserModel
	//global.DB.Debug().Order("age desc").Find(&user) //按照年龄降序
	////global.DB.Debug().Order("age asc").Find(&user) //按照年龄升序
	//fmt.Println(user)

}

func main() {
	global.Connect()
	Highquery()
}
