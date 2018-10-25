package bean

import ()

// UserView UserView
type UserView struct {
	UserID string `gorm:";unique" form:"userID" json:"userID"` // 手机号码
	Name   string `gorm:"not null" form:"name" json:"name"`
	Email  string `gorm:"not null" form:"email" json:"email"`
}

// AdminUserView AdminUserView
type AdminUserView struct {
	UserID string `gorm:";unique" form:"userID" json:"userID"` // 手机号码
	Name   string `gorm:"not null" form:"name" json:"name"`
	Email  string `gorm:"not null" form:"email" json:"email"`
}
