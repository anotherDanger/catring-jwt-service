package repository

import (
	"catering-jwt-service/domain"
	"context"
	"database/sql"
)

type Repository interface {
	Login(ctx context.Context, tx *sql.Tx, entity *domain.Admin) (*domain.Admin, error)
}
