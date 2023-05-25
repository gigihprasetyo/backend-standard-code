package cityhdl

import (
	"context"

	"github.com/gigihprasetyo/backend-standard-code/internal/core/domain"
	"github.com/gigihprasetyo/backend-standard-code/internal/core/ports"
	"github.com/gigihprasetyo/backend-standard-code/internal/core/proto"
	"google.golang.org/grpc"
)

type cityGrpc struct {
	app         *grpc.Server
	cityService ports.CityService
	proto.UnimplementedMyServiceServer
}

func NewCityGrpc(app *grpc.Server, cityService ports.CityService) {
	grpc := &cityGrpc{
		app:         app,
		cityService: cityService,
	}

	proto.RegisterMyServiceServer(app, grpc)
}

func (h *cityGrpc) GetCityList(c context.Context, req *proto.Request) (*proto.Response, error) {
	city := domain.CityParams{
		Name: "Malang",
	}
	_, _, err := h.cityService.GetAll(c, &city)

	if err != nil {
		return nil, err
	}

	return &proto.Response{
		Reply: city.Name,
	}, nil
}
