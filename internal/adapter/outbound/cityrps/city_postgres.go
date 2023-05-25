package cityrps

import (
	"context"
	"errors"

	"github.com/gigihprasetyo/backend-standard-code/internal/core/domain"
	"github.com/gigihprasetyo/backend-standard-code/internal/core/ports"
	"github.com/gigihprasetyo/backend-standard-code/internal/utils/paging"
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"gorm.io/gorm"
)

type cityPostgres struct {
	postgres *gorm.DB
}

func NewCityPostgres(pg *gorm.DB) ports.CityRepository {
	return &cityPostgres{
		postgres: pg,
	}
}

func (instance *cityPostgres) GetAll(ctx context.Context, params *domain.CityParams) ([]*domain.City, *paginator.Cursor, error) {
	var (
		cityModel []*domain.City
	)

	// set default limit paging
	if params.PagingQuery == nil {
		params.PagingQuery = &domain.PagingQuery{}
	}
	if params.PagingQuery.Limit == nil {
		params.PagingQuery.Limit = &paging.DefaultLimit
	}

	q := instance.postgres.Preload("Province")

	// filter by name
	if params.Name != "" {
		q = q.Where("LOWER(name) LIKE LOWER(?)", "%"+params.Name+"%")
	}
	// filter by province id
	if params.ProvinceID != nil {
		q = q.Where("province_id = ?", params.ProvinceID)
	}

	p := paging.Resolver(*params.PagingQuery)

	// set cursor order rule
	p.SetRules(
		paginator.Rule{
			Key:     "Name",
			Order:   paginator.Order(paginator.ASC),
			SQLRepr: "cities.name",
		},
	)

	result, cursor, err := p.Paginate(q, &cityModel)
	if err != nil {
		return nil, nil, err
	}

	if result.Error != nil {
		return nil, nil, err
	}

	return cityModel, &cursor, nil
}

func (instance *cityPostgres) GetDetail(ctx context.Context, params *domain.CityParams) (*domain.City, error) {
	var (
		city *domain.City
	)
	result := instance.postgres.Preload("Province").Where("id = ?", params.ID).First(&city)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return city, nil
}

func (instance *cityPostgres) GetAllByName(ctx context.Context, name string) ([]*domain.City, error) {
	var cities []*domain.City
	result := instance.postgres.Preload("Province").
		Where("LOWER(name) LIKE LOWER(?)", "%"+name+"%").
		Find(&cities)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return cities, nil
}

func (instance *cityPostgres) GetAllWithoutPaginate(ctx context.Context, params *domain.CityParams) ([]*domain.City, error) {
	var (
		cityModel []*domain.City
	)

	q := instance.postgres.Preload("Province")

	// filter by name
	if params.Name != "" {
		q = q.Where("LOWER(name) LIKE LOWER(?)", "%"+params.Name+"%")
	}
	// filter by province id
	if params.ProvinceID != nil {
		q = q.Where("province_id = ?", params.ProvinceID)
	}

	err := q.Find(&cityModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return cityModel, nil
}
