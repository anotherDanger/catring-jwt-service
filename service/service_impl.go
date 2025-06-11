package service

import (
	"catering-jwt-service/domain"
	"context"
	"fmt"
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

func (svc *ServiceImpl) Refresh(ctx context.Context, tokenStr string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	secret := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return err.Error(), err
			return nil, err
		}

		return secret, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("error parsing claims")
		return err.Error(), err
	}

	id, _ := claims["id"].(string)
	username, _ := claims["username"].(string)

	NewClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      time.Now().Add(1 * time.Minute).Unix(),
	})

	newToken, err := NewClaims.SignedString(secret)
	if err != nil {
		return err.Error(), err
		panic(err)
	}

	return newToken, nil
}
