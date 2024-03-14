package jwtn

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("your_secret_key")

// Claims структура для JWT токена
type Claims struct {
	UserID int64  `json:"userId"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// GenerateToken генерирует новый JWT токен для пользователя
func GenerateToken(userID int64, role string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// VerifyToken проверяет JWT токен и извлекает ID пользователя и его роль
func VerifyToken(tokenString string) (int64, string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return 0, "", err
	}

	return claims.UserID, claims.Role, nil
}
