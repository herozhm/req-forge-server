package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// secret 生产环境必须从环境变量读取(如 os.Getenv("JWT_SECRET")),
// 这里用固定值仅为本地演示。
var secret = []byte("req-forge-dev-secret")

func Sign(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(secret)
}

func Verify(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("token 无效")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("token 无效")
	}
	return claims, nil
}
