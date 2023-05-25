package citysvc

import (
	"context"
	"errors"

	"github.com/gigihprasetyo/backend-standard-code/internal/core/domain"
	"github.com/gigihprasetyo/backend-standard-code/internal/core/ports"
	responseErr "github.com/gigihprasetyo/backend-standard-code/internal/error"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

var ErrCityNotFound = errors.New("city not found")
var ErrCityAll = errors.New("failed to get all city")

type cityService struct {
	log      *zap.Logger
	cityRepo ports.CityRepository
}

func NewCityService(log *zap.Logger, cityRepo ports.CityRepository) ports.CityService {
	return &cityService{
		log:      log,
		cityRepo: cityRepo,
	}
}

func (instance *cityService) GetAll(ctx context.Context, params *domain.CityParams) ([]*domain.CityTransformer, *domain.CursorTransform, error) {
	data, cursor, err := instance.cityRepo.GetAll(ctx, params)
	if err != nil {
		instance.log.Error("Failed to get cities: ", zap.Error(err))
		return nil, nil, responseErr.New(fiber.StatusInternalServerError, responseErr.WithMessage(ErrCityAll.Error()))
	}

	if data == nil {
		return nil, nil, nil
	}

	cursorTrans := domain.CursorTransform{
		After:  cursor.After,
		Before: cursor.Before,
	}

	var results []*domain.CityTransformer
	for _, data := range data {
		results = append(results, &domain.CityTransformer{
			ID:   data.ID,
			Name: data.Name,
		})
	}

	return results, &cursorTrans, nil
}

func (instance *cityService) GetDetail(ctx context.Context, params *domain.CityParams) (*domain.CityTransformer, error) {
	getCity, err := instance.cityRepo.GetDetail(ctx, params)
	if err != nil {
		instance.log.Error("Failed to get detail city: ", zap.Error(err))
		return nil, responseErr.New(fiber.StatusInternalServerError, responseErr.WithMessage(ErrCityNotFound.Error()))
	}

	if getCity == nil {
		return nil, nil
	}

	result := &domain.CityTransformer{
		ID:   getCity.ID,
		Name: getCity.Name,
	}

	return result, nil
}
