package repo

import (
	"context"

	"github.com/andresxlp/go-echo-postgres-hexa/internal/domain/ports/repo"
	"github.com/andresxlp/go-echo-postgres-hexa/internal/infra/adapters/postgresql/models"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repo.UserRepositoryInterface {
	return &userRepository{db}
}

func (ur *userRepository) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	return users, nil
}
