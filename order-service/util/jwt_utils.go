package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strings"
	"time"
)

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

const defaultJWTSecret = "order-ez-development-secret"

func jwtSigningKey() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = defaultJWTSecret
	}
	return []byte(secret)
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
	tokenString, err := token.SignedString(jwtSigningKey())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

var (
	ErrInvalidToken         = errors.New("无效的令牌")
	ErrExpiredToken         = errors.New("令牌已过期")
	ErrInvalidSigningMethod = errors.New("无效的签名方法")
)

func ValidateToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}
		return jwtSigningKey(), nil
	})

	if err != nil {
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

func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPasswordBytes), nil
}

func ParseBearerToken(authHeader string) (string, error) {
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("无效的 Authorization 格式，应为 Bearer <token>")
	}
	return parts[1], nil
}
