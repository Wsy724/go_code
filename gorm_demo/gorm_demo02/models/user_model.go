package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	Id              int
	Name            string           `gorm:"size:16; not null; unique"`
	Age             int              `gorm:"default:18"`
	UserDetailModel *UserDetailModel `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` //constraint:OnDelete:CASCADE级联删除
	CreatedAt       time.Time
	//DeletedAt gorm.DeletedAt //软删除
}

type UserDetailModel struct {
	ID        int
	UserID    int        `gorm:"unique"` //一对一的情况下需要加上唯一约束
	UserModel *UserModel `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Email     string     `gorm:"size:32"`
}

// 钩子函数
func (u *UserModel) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("创建的钩子函数")
	return nil
}

func (u *UserModel) BeforeUpdate(tx *gorm.DB) error {
	fmt.Println("更新的的钩子函数")
	return nil
}

func (u *UserModel) BeforeDelete(tx *gorm.DB) error {
	fmt.Println("删除的钩子函数")
	return nil
}
