package ports

import (
	"context"

	"github.com/gigihprasetyo/backend-standard-code/internal/core/domain"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
)

type (
	CityRepository interface {
		GetAll(ctx context.Context, params *domain.CityParams) ([]*domain.City, *paginator.Cursor, error)
		GetDetail(ctx context.Context, params *domain.CityParams) (*domain.City, error)
		GetAllByName(ctx context.Context, name string) ([]*domain.City, error)
		GetAllWithoutPaginate(ctx context.Context, params *domain.CityParams) ([]*domain.City, error)
	}
)
