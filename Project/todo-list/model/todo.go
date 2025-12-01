package model

import (
	"time"

	"gorm.io/gorm"
)

// Todo 待办事项模型
type Todo struct {
	ID        uint           `gorm:"primarykey" json:"id"`           // 主键ID
	Title     string         `gorm:"size:255;not null" json:"title"` // 待办标题
	Content   string         `gorm:"type:text" json:"content"`       // 待办内容
	Status    int            `gorm:"default:0" json:"status"`        // 状态：0-未完成 1-已完成
	CreatedAt time.Time      `json:"created_at"`                     // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`                     // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                 // 软删除（不显示在JSON）
}

// 自动迁移表结构（启动时执行）
func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&Todo{})
	if err != nil {
		panic("表迁移失败：" + err.Error())
	}
}
