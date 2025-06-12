package repository

import (
	"catering-jwt-service/domain"
	"context"
	"database/sql"
)

type RepositoryImpl struct{}

func NewRepositoryImpl() Repository {
	return &RepositoryImpl{}
}

func (repo *RepositoryImpl) Login(ctx context.Context, tx *sql.Tx, entity *domain.Admin) (*domain.Admin, error) {
	query := "select id, username, password from admin where username = ?"
	result := tx.QueryRowContext(ctx, query, entity.Username)

	var response domain.Admin
	result.Scan(&response.Id, &response.Username, &response.Password)

	return &response, nil

}
