package main

import (
	"gorm_demo02/global"
)

func oneToOne() {
	//创建用户，连带着创建用户详情
	//err := global.DB.Create(&models.UserModel{
	//	Name: "唐僧",
	//	Age:  18,
	//	UserDetailModel: &models.UserDetailModel{
	//		Email: "xxxx@qq.com",
	//	},
	//}).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//创建用户详情，并关联用户
	//err := global.DB.Create(&models.UserDetailModel{
	//	Email:     "143242@qq.com",
	//	UserModel: &models.UserModel{Id: 14},
	//	//UserID: 14
	//}).Error
	//if err != nil {
	//	fmt.Println(err)
	//}

	//通过用户详情查用户 （正查）
	//var id = 14
	//var detail models.UserDetailModel
	//global.DB.Preload("UserModel").Take(&detail, "user_id = ?", id)
	//fmt.Println(detail.Email, detail.UserModel.Name)

	//通过用户查用户详情（反查）
	//var id = 14
	//var user models.UserModel
	//global.DB.Preload("UserDetailModel").Take(&user, id)
	//fmt.Println(user.Name, user.UserDetailModel.Email)

	//级联删除（删除用户，用户详情也被删除）
	//var user models.UserModel
	//global.DB.Take(&user, 13)
	//global.DB.Select("UserDetailModel").Delete(&user)

	//constraint:OnDelete:CASCADE;使用这个后可以直接删除，关联的那个也会自动删除(前提是生成实体外键)
	//var user models.UserModel
	//global.DB.Take(&user, 2)
	//global.DB.Delete(&user)
}

func main() {
	global.Connect()
	//global.Migrate()
	//oneToOne()
}
