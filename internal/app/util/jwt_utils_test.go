package util

import (
	"fmt"
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	key, _ := generateSecretKey()
	fmt.Println("-----???---------")
	fmt.Println(key)
	fmt.Println("------????--------")

	token, _ := GenerateToken("z")
	fmt.Println("-----???---------")
	fmt.Println(token)
	fmt.Println("------????--------")

	validateToken, _ := ValidateToken(token)
	fmt.Println("-----???---------")
	fmt.Printf("%+v", validateToken)
	fmt.Println("------????--------")
	//// 定义一个测试用的用户 ID
	//userID := uint(123)
	//
	//// 生成令牌
	//tokenString, err := GenerateToken(userID)
	//if err != nil {
	//	t.Fatalf("GenerateToken failed: %v", err)
	//}
	//
	//// 验证令牌
	//claims, err := ValidateToken(tokenString)
	//if err != nil {
	//	t.Fatalf("ValidateToken failed: %v", err)
	//}
	//
	//// 检查用户 ID 是否匹配
	//if claims.UserClaimsStr != userID {
	//	t.Errorf("Expected user ID %d, but got %d", userID, claims.UserClaimsStr)
	//}
	//
	//// 检查令牌是否在有效期内
	//now := time.Now().Unix()
	//if claims.ExpiresAt < now {
	//	t.Errorf("Token has expired. Expiration time: %d, Current time: %d", claims.ExpiresAt, now)
	//}
}
