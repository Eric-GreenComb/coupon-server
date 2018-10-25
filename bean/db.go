package bean

import (
	"github.com/jinzhu/gorm"
)

// Users Users
type Users struct {
	gorm.Model
	User
}

// User User
type User struct {
	UserID string `gorm:";unique" form:"userID" json:"userID"` // 手机号码
	Name   string `gorm:"not null" form:"name" json:"name"`
	Passwd string `gorm:"not null" form:"passwd" json:"passwd"`
	Email  string `gorm:"not null" form:"email" json:"email"`
}

// AdminUsers AdminUsers
type AdminUsers struct {
	gorm.Model
	User
}
