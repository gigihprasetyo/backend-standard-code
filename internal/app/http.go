package app

import (
	"github.com/go-redis/redis/v8"
	"github.com/pasarin-tech/pasarin-core/internal/adapter/inbound/cityhdl"
	"github.com/pasarin-tech/pasarin-core/internal/adapter/outbound/cityrps"
	"github.com/pasarin-tech/pasarin-core/internal/core/services/citysvc"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handlers struct {
	Postgres *gorm.DB
	R        *fiber.App
	Logger   *zap.Logger
	//ElasticSearch *elastic.Client
	Redis *redis.Client
}

func (h *Handlers) SetupRouter() {
	h.R.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("OK")
	})
	h.R.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Talk Is Cheap Show Me Your Code")
	})
	//initialize Repository
	//productElasticSearchRep := productrps.NewProductElasticSearch(h.ElasticSearch)

	cityRep := cityrps.NewCityPostgres(h.Postgres)

	//initialize bussiness
	cityService := citysvc.NewCityService(h.Logger, cityRep)

	//handlers initialize
	cityhdl.NewCityHandler(h.R, cityService)

}
