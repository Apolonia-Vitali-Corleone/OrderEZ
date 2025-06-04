package repository

import (
	"OrderEZ/internal/app/model"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) Save(user model.User) error {
	// 对密码进行加密
	passwordBytes := []byte(user.Password)
	password, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(password)
	return r.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Save(&user)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}

// GetAllUsers 方法使用 GORM 查询指定页码和每页数量的用户
func (r *UserRepository) GetAllUsers(page, pageSize int) ([]model.User, error) {
	var users []model.User
	offset := (page - 1) * pageSize
	result := r.db.Offset(offset).Limit(pageSize).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
