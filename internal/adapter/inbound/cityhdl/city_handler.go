package cityhdl

import (
	"strconv"

	"github.com/pasarin-tech/pasarin-core/internal/core/domain"
	"github.com/pasarin-tech/pasarin-core/internal/core/ports"
	"github.com/pasarin-tech/pasarin-core/internal/core/services/citysvc"
	"github.com/pasarin-tech/pasarin-core/internal/response"

	responseErr "github.com/pasarin-tech/pasarin-core/internal/error"

	"github.com/gofiber/fiber/v2"
)

type cityHandler struct {
	app         *fiber.App
	cityService ports.CityService
}

func NewCityHandler(app *fiber.App, cityService ports.CityService) {
	cityHdl := cityHandler{
		app:         app,
		cityService: cityService,
	}
	pApi := cityHdl.app.Group("/v1/cities")
	pApi.Get("/", cityHdl.getAll)
	pApi.Get("/:id", cityHdl.getDetail)
}

func (instance *cityHandler) getAll(c *fiber.Ctx) error {
	requestParams := new(domain.CityParams)
	if err := c.QueryParser(requestParams); err != nil {
		return responseErr.Response(c, responseErr.New(fiber.StatusBadRequest, responseErr.WithMessage(responseErr.ErrBadRequest.Error())))
	}

	city, cursor, err := instance.cityService.GetAll(c.Context(), requestParams)
	if err != nil {
		return responseErr.Response(c, responseErr.New(fiber.StatusInternalServerError, responseErr.WithMessage(citysvc.ErrCityAll.Error())))
	}
	return response.Success(c, fiber.StatusOK, response.SuccessData(city), response.SuccessMeta(cursor))
}

func (instance *cityHandler) getDetail(c *fiber.Ctx) error {
	cityID, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return responseErr.Response(c, responseErr.New(fiber.StatusBadRequest, responseErr.WithMessage(responseErr.ErrBadRequest.Error())))
	}

	requestParams := &domain.CityParams{
		ID: cityID,
	}

	city, err := instance.cityService.GetDetail(c.Context(), requestParams)
	if err != nil {
		return responseErr.Response(c, responseErr.New(fiber.StatusNotFound, responseErr.WithMessage(citysvc.ErrCityNotFound.Error())))
	}
	return response.Success(c, fiber.StatusOK, response.SuccessData(city))
}
