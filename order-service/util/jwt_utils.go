package util

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type Claims struct {
	UserID int64 `json:"user_id"`
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

func GenerateToken(userID int64) (string, error) {
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

// 定义自定义错误类型
var (
	ErrInvalidToken         = errors.New("无效的令牌")
	ErrExpiredToken         = errors.New("令牌已过期")
	ErrInvalidSigningMethod = errors.New("无效的签名方法")
)

func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法是否正确（防止算法替换攻击）
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return jwtKey, nil
	})

	if err != nil {
		// 区分不同错误类型（如过期、签名错误等）
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrExpiredToken
			}
		}
		return nil, ErrInvalidToken
	}

	if !token.Valid {
		return nil, ErrInvalidToken
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

// ParseBearerToken 辅助函数：解析 Bearer Token
func ParseBearerToken(authHeader string) (string, error) {
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("无效的 Authorization 格式，应为 Bearer <token>")
	}
	return parts[1], nil
}
