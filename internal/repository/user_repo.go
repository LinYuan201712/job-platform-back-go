package repository

import (
	"gorm.io/gorm"
	"job-platform-go/internal/model/entity"
	"job-platform-go/pkg/database"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 构造函数
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.DB,
	}
}

// FindByEmail 根据邮箱查找用户
func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUserWithRole
func (r *UserRepository) CreateUserWithRole(user *entity.User, roleStr string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		//1.插入User表
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		//2.根据角色插入关联表
		if roleStr == "student" {
			student := entity.Student{
				UserID: user.ID,
			}
			if err := tx.Create(&student).Error; err != nil {
				return err
			}

		} else if roleStr == "hr" {
			company := entity.Company{
				UserID: user.ID,
			}
			if err := tx.Create(&company).Error; err != nil {
				return err
			}
		}
		return nil
	})

}

// UpdateLastLogin 更新最后登录时间
func (r *UserRepository) UpdateLastLogin(UserID int) {
	now := time.Now()
	r.db.Model(&entity.User{}).Where("id=?", UserID).Update("last_login_at", now)
}
