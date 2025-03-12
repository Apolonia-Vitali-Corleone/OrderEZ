package util

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// 生成并初始化 jwtKey
var jwtKey []byte

func init() {
	key, err := generateSecretKey()
	if err != nil {
		panic("Failed to generate secret key: " + err.Error())
	}
	jwtKey = []byte(key)
}

func generateSecretKey() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic("创建私钥失败！")
	}
	return base64.StdEncoding.EncodeToString(key), nil
}

func GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken 解析并验证 JWT 令牌
func ValidateToken(tokenString string) (*Claims, error) {
	// 创建了一个内容为空？零值？的Claims实例
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// HashPassword 加密密码
func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPasswordBytes), nil
}
