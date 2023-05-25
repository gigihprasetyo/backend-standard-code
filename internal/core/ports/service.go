package ports

import (
	"context"

	"github.com/gigihprasetyo/backend-standard-code/internal/core/domain"
	//"github.com/gigihprasetyo/backend-standard-code/internal/core/middleware"
)

type (
	CityService interface {
		GetAll(ctx context.Context, params *domain.CityParams) ([]*domain.CityTransformer, *domain.CursorTransform, error)
		GetDetail(ctx context.Context, params *domain.CityParams) (*domain.CityTransformer, error)
	}
)
