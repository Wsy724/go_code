package main

import (
	"fmt"
	"gorm_demo02/global"
	"gorm_demo02/models"
)

func oneToMany() {
	//创建：
	//1.来了个女神自带两个舔狗
	//err := global.DB.Create(&models.GirlModel{
	//	Name: "小芳",
	//	BoyList: []models.BoyModel{
	//		{Name: "一号"},
	//		{Name: "二号"},
	//	},
	//}).Error
	//fmt.Println(err)

	//2.来了个女神选了个舔狗
	//err := global.DB.Create(&models.GirlModel{
	//	Name: "小爱",
	//	BoyList: []models.BoyModel{
	//		{ID: 2},
	//	},
	//}).Error
	//fmt.Println(err)

	//3.来了个舔狗选了个女神
	//err := global.DB.Create(&models.BoyModel{
	//	Name:   "三号",
	//	GirlId: 2,
	//}).Error
	//fmt.Println(err)

	//查询小爱的舔狗
	var girl models.GirlModel
	global.DB.Preload("BoyList").Take(&girl, "name=?", "小爱")
	fmt.Println(girl.Name, girl.BoyList, len(girl.BoyList))

	girl = models.GirlModel{}
	global.DB.Preload("BoyList", "name=?", "三号").Take(&girl, "name=?", "小爱")
	fmt.Println(girl.Name, girl.BoyList, len(girl.BoyList))

	//专门查关联
	girl = models.GirlModel{}
	global.DB.Take(&girl, "name=?", "小爱")
	var boyList []models.BoyModel
	global.DB.Model(&girl).Association("BoyList").Find(&boyList)
	fmt.Println(girl.Name, boyList)
}

func main() {
	global.Connect()
	//global.Migrate()
	oneToMany()

}
