package repo

import (
	"context"

	"github.com/andresxlp/go-echo-postgres-hexa/internal/infra/adapters/postgresql/models"
)

type UserRepositoryInterface interface {
	GetAll(ctx context.Context) ([]models.User, error)
}
