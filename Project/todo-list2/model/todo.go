package model

import "gorm.io/gorm"

//Todo 待办事项模型
type Todo struct {
	gorm.Model
	Title   string `gorm:"size:255;not null" json:"title"` //代办标题
	Content string `gorm:"type:text" json:"content"`       //代办内容
	Status  int    `gorm:"default:0" json:"status"`        // 状态：0-未删除 1-已完成
}

//自动迁移表结构(启动时执行)
func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&Todo{})
	if err != nil {
		panic("表迁移失败：" + err.Error())
	}
}
