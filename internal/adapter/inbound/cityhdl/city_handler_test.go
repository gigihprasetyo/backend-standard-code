package cityhdl_test

import (
	"testing"
	"time"

	"github.com/gigihprasetyo/backend-standard-code/internal/core/domain"
	"github.com/gigihprasetyo/backend-standard-code/internal/core/ports/mocks"
	"github.com/gigihprasetyo/backend-standard-code/internal/core/services/citysvc"
	responseErr "github.com/gigihprasetyo/backend-standard-code/internal/error"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type CityTestSuite struct {
	suite.Suite
	log             *zap.Logger
	City            domain.City
	MockCityService *mocks.CityService
}

func (suite *CityTestSuite) SetupTest() {
	var (
		cityName string = "KOTA MALANG"
	)
	suite.log, _ = zap.NewDevelopment()

	// suite model
	suite.City = domain.City{
		ID:        3573,
		Name:      cityName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.MockCityService = new(mocks.CityService)
}

func (suite *CityTestSuite) TestGetAll() {
	//app := fiber.New()
	cases := []struct {
		name            string
		description     string
		err             error
		cityService     *mocks.CityService
		cityTransformer []*domain.CityTransformer
		paramID         string
		pagination      *domain.CursorTransform
		request         map[string]interface{}
		response        map[string]interface{}
	}{
		{
			name:            "get_cities_failed",
			description:     "get cities failed",
			err:             responseErr.New(fiber.StatusInternalServerError, responseErr.WithMessage(citysvc.ErrCityAll.Error())),
			cityService:     suite.MockCityService,
			cityTransformer: []*domain.CityTransformer{},
			request:         nil,
			response: map[string]interface{}{
				"message": citysvc.ErrCityAll.Error(),
			},
		},
		{
			name:        "get_cities_success",
			description: "get cities success",
			err:         nil,
			cityService: suite.MockCityService,
			cityTransformer: []*domain.CityTransformer{
				suite.City.ToTransformer(),
			},
			pagination: &domain.CursorTransform{
				After:  nil,
				Before: nil,
			},
			request: nil,
			response: map[string]interface{}{
				"data": []*domain.CityTransformer{
					suite.City.ToTransformer(),
				},
				"meta": &domain.CursorTransform{
					After:  nil,
					Before: nil,
				},
			},
		},
	}

	for _, tCase := range cases {
		switch tCase.name {
		// 	case "get_cities_failed":
		// 		suite.Run(tCase.name, func() {
		// 			cityhdl.NewCityHandler(app, tCase.cityService)
		// 			tCase.cityService.On("GetAll", mock.Anything, mock.AnythingOfType("*domain.CityParams")).Return(tCase.cityTransformer, tCase.pagination, tCase.err).Once()
		// 			encoded, err := json.Marshal(nil)
		// 			if err != nil {

		// 			}
		// 			req := httptest.NewRequest("GET", "/v1/cities", bytes.NewReader(encoded))
		// 			req.Header.Add("Content-Type", "application/json")
		// 			res, err := app.Test(req)
		// 			if err != nil {
		// 			}
		// 			result, err := ioutil.ReadAll(res.Body)
		// 			if err != nil {
		// 			}
		// 			defer res.Body.Close()
		// 			expected, err := json.Marshal(tCase.response)
		// 			if err != nil {
		// 			}
		// 		})
		// 	case "get_cities_success":
		// 		suite.Run(tCase.name, func() {
		// 			cityhdl.NewCityHandler(app, tCase.cityService)
		// 			tCase.cityService.On("GetAll", mock.Anything, mock.AnythingOfType("*domain.CityParams")).Return(tCase.cityTransformer, tCase.pagination, tCase.err).Once()
		// 			encoded, err := json.Marshal(nil)
		// 			if err != nil {
		// 				suite.FailNow(err.Error())
		// 			}
		// 			req := httptest.NewRequest("GET", "/v1/cities", bytes.NewReader(encoded))
		// 			req.Header.Add("Content-Type", "application/json")
		// 			res, err := app.Test(req)
		// 			if err != nil {
		// 				suite.FailNow(err.Error())
		// 			}
		// 			result, err := ioutil.ReadAll(res.Body)
		// 			if err != nil {
		// 				suite.FailNow(err.Error())
		// 			}
		// 			defer res.Body.Close()
		// 			expected, err := json.Marshal(tCase.response)
		// 			if err != nil {
		// 				suite.FailNow(err.Error())
		// 			}

		// 			assert.Equal(suite.T(), fiber.StatusOK, res.StatusCode)
		// 			assert.Equal(suite.T(), expected, result)
		// 		})

		}
	}
}

func (suite *CityTestSuite) TestGetDetail() {
	//app := fiber.New()
	cases := []struct {
		name            string
		description     string
		err             error
		cityService     *mocks.CityService
		cityTransformer *domain.CityTransformer
		params          *domain.CityParams
		paramID         string
		request         map[string]interface{}
		response        map[string]interface{}
	}{
		{
			name:            "param_parse_failed",
			description:     "when param parse failed",
			err:             responseErr.New(fiber.StatusBadRequest, responseErr.WithMessage(responseErr.ErrBadRequest.Error())),
			cityService:     suite.MockCityService,
			cityTransformer: &domain.CityTransformer{},
			paramID:         "ASDasd",
			params: &domain.CityParams{
				ID: 1,
			},
			request: nil,
			response: map[string]interface{}{
				"message": responseErr.ErrBadRequest.Error(),
			},
		},
		{
			name:            "get_city_detail_failed",
			description:     "get city detail failed",
			err:             responseErr.New(fiber.StatusInternalServerError, responseErr.WithMessage(citysvc.ErrCityNotFound.Error())),
			cityService:     suite.MockCityService,
			cityTransformer: &domain.CityTransformer{},
			params: &domain.CityParams{
				ID: 1,
			},
			request: nil,
			response: map[string]interface{}{
				"message": citysvc.ErrCityNotFound.Error(),
			},
		},
		{
			name:            "get_city_detail_success",
			description:     "get city detail success",
			err:             nil,
			cityService:     suite.MockCityService,
			cityTransformer: suite.City.ToTransformer(),
			params: &domain.CityParams{
				ID: 11,
			},
			request: nil,
			response: map[string]interface{}{
				"data": suite.City.ToTransformer(),
			},
		},
	}

	for _, tCase := range cases {
		switch tCase.name {
		// 	case "param_parse_failed":
		// 		suite.Run(tCase.name, func() {
		// 			cityhdl.NewCityHandler(app, tCase.cityService)
		// 			tCase.cityService.On("GetDetail", mock.Anything, tCase.params).Return(tCase.cityTransformer, tCase.err).Once()
		// 			encoded, err := json.Marshal(nil)
		// 			if err != nil {
		// 				suite.FailNow(err.Error())
		// 			}
		// 			url := fmt.Sprintf("/v1/cities/%s", tCase.paramID)
		// 			req := httptest.NewRequest("GET", url, bytes.NewReader(encoded))
		// 			req.Header.Add("Content-Type", "application/json")
		// 			res, err := app.Test(req)
		// 			if err != nil {
		// 				suite.FailNow(err.Error())
		// 			}
		// 			_, err = ioutil.ReadAll(res.Body)
		// 			if err != nil {
		// 				suite.FailNow(err.Error())
		// 			}
		// 			defer res.Body.Close()
		// 			_, err = json.Marshal(tCase.response)
		// 			if err != nil {
		// 				suite.FailNow(err.Error())
		// 			}

		// 			_, err = strconv.ParseUint(tCase.paramID, 10, 64)
		// 			if err != nil {
		// 				assert.Equal(suite.T(), fiber.StatusBadRequest, res.StatusCode)
		// 			}

		// 		})
		// 	case "get_city_detail_failed":
		// 		suite.Run(tCase.name, func() {
		// 			cityhdl.NewCityHandler(app, tCase.cityService)
		// 			tCase.cityService.On("GetDetail", mock.Anything, tCase.params).Return(tCase.cityTransformer, tCase.err).Once()
		// 			encoded, err := json.Marshal(nil)
		// 			if err != nil {
		// 				suite.FailNow(err.Error())
		// 			}
		// 			url := fmt.Sprintf("/v1/cities/%d", tCase.params.ID)
		// 			req := httptest.NewRequest("GET", url, bytes.NewReader(encoded))
		// 			req.Header.Add("Content-Type", "application/json")
		// 			res, err := app.Test(req)
		// 			if err != nil {
		// 				suite.FailNow(err.Error())
		// 			}
		// 			result, err := ioutil.ReadAll(res.Body)
		// 			if err != nil {
		// 				suite.FailNow(err.Error())
		// 			}
		// 			defer res.Body.Close()
		// 			expected, err := json.Marshal(tCase.response)
		// 			if err != nil {
		// 				suite.FailNow(err.Error())
		// 			}

		// 			assert.Equal(suite.T(), fiber.StatusNotFound, res.StatusCode)
		// 			assert.Equal(suite.T(), expected, result)
		// 		})
		// 	case "get_city_detail_success":
		// 		suite.Run(tCase.name, func() {
		// 			cityhdl.NewCityHandler(app, tCase.cityService)
		// 			tCase.cityService.On("GetDetail", mock.Anything, tCase.params).Return(tCase.cityTransformer, tCase.err).Once()
		// 			encoded, err := json.Marshal(nil)
		// 			if err != nil {
		// 				suite.FailNow(err.Error())
		// 			}
		// 			url := fmt.Sprintf("/v1/cities/%d", tCase.params.ID)
		// 			req := httptest.NewRequest("GET", url, bytes.NewReader(encoded))
		// 			req.Header.Add("Content-Type", "application/json")
		// 			res, err := app.Test(req)
		// 			if err != nil {
		// 				suite.FailNow(err.Error())
		// 			}
		// 			result, err := ioutil.ReadAll(res.Body)
		// 			if err != nil {
		// 				suite.FailNow(err.Error())
		// 			}
		// 			defer res.Body.Close()
		// 			expected, err := json.Marshal(tCase.response)
		// 			if err != nil {
		// 				suite.FailNow(err.Error())
		// 			}

		// 			assert.Equal(suite.T(), fiber.StatusOK, res.StatusCode)
		// 			assert.JSONEq(suite.T(), string(expected), string(result))
		// 		})
		}
	}
}

func TestCitySuite(t *testing.T) {
	suite.Run(t, new(CityTestSuite))
}
