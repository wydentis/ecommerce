package products

import (
	"context"

	repository "github.com/wydentis/ecommerce/internal/adapters/postgresql/sqlc"
)

type Service interface {
	ListProducts(ctx context.Context) ([]repository.Product, error)
}

type svc struct {
	repo repository.Querier
}

func NewService(repo repository.Querier) Service {
	return &svc{repo: repo}
}

func (s *svc) ListProducts(ctx context.Context) ([]repository.Product, error) {
	return s.repo.ListProducts(ctx)
}
