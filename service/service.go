package service

import (
	"catering-jwt-service/domain"
	"context"
)

type Service interface {
	// Login(ctx context.Context, request *domain.Admin) (*web.Response, error)
	Register(ctx context.Context, entity *domain.Domain) (string, error)
	Refresh(ctx context.Context, tokenStr string) (token string, username string, err error)
}
