package ports

import (
	"context"

	"github.com/pasarin-tech/pasarin-core/internal/core/domain"
	"github.com/pasarin-tech/pasarin-core/internal/core/middleware"
)

type (
	CityService interface {
		GetAll(ctx context.Context, params *domain.CityParams) ([]*domain.CityTransformer, *domain.CursorTransform, error)
		GetDetail(ctx context.Context, params *domain.CityParams) (*domain.CityTransformer, error)
	}
)
