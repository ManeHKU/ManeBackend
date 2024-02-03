package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	secretKey string
}

type UserClaims struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
}

func NewJWTManager(secretKey string) *Manager {
	return &Manager{secretKey}
}

func (manager *Manager) VerifyUserToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.secretKey), nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*UserClaims); ok {
		return claims, nil
	}

	return nil, errors.New("claim failed")
}
