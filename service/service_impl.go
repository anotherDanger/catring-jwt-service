package service

import (
	"catering-jwt-service/domain"
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type ServiceImpl struct {
}

func NewServiceImpl() Service {
	return &ServiceImpl{}
}

func (svc *ServiceImpl) Register(ctx context.Context, entity *domain.Domain) (string, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       entity.Id,
		"username": entity.Username,
		"exp":      time.Now().Add(1 * time.Minute).Unix(),
	})

	tokenT := os.Getenv("JWT_SECRET")
	byteToken := []byte(tokenT)

	token, err := claims.SignedString(byteToken)
	if err != nil {
		panic(err)
	}

	return token, nil
}

func (svc *ServiceImpl) RefreshToken(ctx context.Context, entity *domain.Domain) (string, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       entity.Id,
		"username": entity.Username,
		"exp":      time.Now().Add(7 * 24 * time.Hour).Unix(),
	})

	tokenT := os.Getenv("JWT_SECRET")
	byteToken := []byte(tokenT)

	token, err := claims.SignedString(byteToken)
	if err != nil {
		panic(err)
	}

	return token, nil
}

func (svc *ServiceImpl) Refresh(ctx context.Context, tokenStr string) (token string, username string, err error) {
	err = godotenv.Load()
	if err != nil {
		panic(err)
	}
	secret := []byte(os.Getenv("JWT_SECRET"))

	parsedToken, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, errors.New("invalid token")
		}
		return secret, nil
	})

	if err != nil {
		return "", "", err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return "", "", errors.New("invalid token claims")
	}

	id, _ := claims["id"].(string)
	username, _ = claims["username"].(string)

	newClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(1 * time.Minute).Unix(),
	})

	newToken, err := newClaims.SignedString(secret)
	if err != nil {
		return "", "", err
	}

	return newToken, username, nil
}
