package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"job-platform-go/internal/config"
)

type Claims struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT
func GenerateToken(id int, email string, role string) (string, error) {
	jwtConfig := config.GlobalConfig.Security
	claims := Claims{
		ID:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(jwtConfig.ExpirationMs) * time.Millisecond)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtConfig.Secret))
}

// ParseToken 解析 JWT
// 输入 token → 用密钥解密 → 用 Claims 接收内容 → 校验 → 成功就返回用户信息 → 否则返回错误。
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GlobalConfig.Security.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
