package service

import (
	"catering-jwt-service/domain"
	"context"
)

type Service interface {
	Register(ctx context.Context, entity *domain.Domain) (string, error)
	Refresh(ctx context.Context, tokenStr string) (token string, username string, err error)
	RefreshToken(ctx context.Context, entity *domain.Domain) (string, error)
}
