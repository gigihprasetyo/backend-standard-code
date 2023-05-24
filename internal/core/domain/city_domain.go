package domain

import (
	"time"

	"gorm.io/gorm"
)

type City struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	ProvinceID uint64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt

	Province *Province
}

type CityTransformer struct {
	ID                       uint64                   `json:"id"`
	Name                     string                   `json:"name"`
	ProvinceBasicTransformer ProvinceBasicTransformer `json:"province,omitempty"`
}

type CityBasicTransformer struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type CityParams struct {
	ID         uint64
	Name       *string `query:"name"`
	ProvinceID *uint64 `query:"province_id"`
	*PagingQuery
}

func (c *City) ToTransformer() *CityTransformer {
	return &CityTransformer{
		ID:                       c.ID,
		Name:                     c.Name,
		ProvinceBasicTransformer: *c.Province.ToBasicTransformer(),
	}
}

func (c *City) ToBasicTransformer() *CityBasicTransformer {
	return &CityBasicTransformer{
		ID:   c.ID,
		Name: c.Name,
	}
}
