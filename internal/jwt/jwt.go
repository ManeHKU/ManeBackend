package jwt

import (
	"ManeBackend/internal/env"
	"context"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
)

var keyFunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
	config := env.GetConfig()
	return []byte(config.JWT_SECRET), nil
}

func GetUserID(ctx context.Context) string {
	claims, exist := GetJWTClaims(ctx)
	if !exist {
		return ""
	}
	subject, _ := claims.GetSubject()
	return subject
}

func GetJWTClaims(ctx context.Context) (*jwt.MapClaims, bool) {
	metadataMap, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		return nil, false
	}
	authorizationKeys := metadataMap.Get("authorization")
	if len(authorizationKeys) != 1 {
		return nil, false
	}
	tokenString := authorizationKeys[0]
	token, _ := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, keyFunc)

	if token != nil {
		if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
			return claims, true
		}
	}
	return nil, false
}
