package entity

import (
	"time"
)

type UserRole int

const (
	RoleStudent UserRole = 1
	RoleHR      UserRole = 2
)

type UserStatus int

const (
	StatusPending  UserStatus = 0
	StatusActive   UserStatus = 1
	StatusDisabled UserStatus = 2
)

type User struct {
	ID           int        `gorm:"primaryKey"`
	Email        string     `gorm:"unique;not null" json:"email"`
	PasswordHash string     `gorm:"column:password_hash;not null" json:"-"` //不返回给前端
	Role         UserRole   `gorm:"not null" json:"role"`
	Status       UserStatus `gorm:"default:0" json:"status"`
	LastLoginAt  *time.Time `json:"last_login_at"` // 使用指针允许为 null
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

// TableName
func (User) TableName() string {
	return "users"
}
