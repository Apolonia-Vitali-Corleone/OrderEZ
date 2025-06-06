package service

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
	"user-service/infrastructure/messaging"
	"user-service/internal/model/po"
	"user-service/internal/repository"
	"user-service/util"
)

type UserService struct {
	userRepo *repository.UserRepository
	redis    *redis.Client
	rabbitMQ *messaging.RabbitMQ
}

func NewUserService(db *gorm.DB, redisClient *redis.Client, mq *messaging.RabbitMQ) *UserService {
	userRepo := repository.NewUserRepository(db)
	return &UserService{
		userRepo: userRepo,
		redis:    redisClient,
		rabbitMQ: mq,
	}
}

func (s *UserService) Login(username, password string) (string, error) {
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

	token, err := util.GenerateToken(user.UserID)
	if err != nil {
		return "", err
	}

	tokenExpire := time.Minute * 5

	// 将 token 存入 Redis 缓存会话数据
	err = s.redis.Set(context.Background(), "login:token:"+token, user.UserID, tokenExpire).Err()
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) Register(username, password string) (string, error) {
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
	var newUser po.User
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
	token, err := util.GenerateToken(byUsername.UserID)
	if err != nil {
		return "", err
	}

	return token, nil

}

// GetAllUsers 方法用于获取所有用户
func (s *UserService) GetAllUsers(page, pageSize int) ([]po.User, error) {
	return s.userRepo.GetAllUsers(page, pageSize)
}

// Logout 登出
func (s *UserService) Logout(token string) error {
	// 从 Redis 中删除令牌
	err := s.redis.Del(context.Background(), token).Err()
	if err != nil {
		return err
	}
	return nil
}
