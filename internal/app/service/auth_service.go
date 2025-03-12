package service

import (
	"OrderEZ/internal/app/model"
	"OrderEZ/internal/app/repository"
	"OrderEZ/internal/app/util"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type AuthService struct {
	userRepo *repository.UserRepository
	redis    *redis.Client
}

func NewAuthService(redisClient *redis.Client, db *gorm.DB) *AuthService {
	userRepo := repository.NewUserRepository(db)
	return &AuthService{
		userRepo: userRepo,
		redis:    redisClient,
	}
}

func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	token, err := util.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	// 将 token 存入 Redis 缓存会话数据
	err = s.redis.Set(context.Background(), user.Username, token, 24*time.Hour).Err()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) Register(username, password string) (string, error) {
	// 前端进行数据验证，确保我们这里收到的数据都是有效的
	existingUser, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	if existingUser != nil {
		return "", errors.New("username already exists")
	}

	// 加密密码
	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return "", err
	}

	// 创建新用户
	var newUser model.User
	newUser.Username = username
	newUser.Password = hashedPassword

	// 保存用户信息
	err = s.userRepo.Save(newUser)
	if err != nil {
		return "", err
	}

	// 获取到他的id
	byUsername, err := s.userRepo.GetUserByUsername(username)

	// 生成 JWT 令牌
	token, err := util.GenerateToken(byUsername.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}
