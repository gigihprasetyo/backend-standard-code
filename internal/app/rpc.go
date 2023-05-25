package app

import (
	"github.com/gigihprasetyo/backend-standard-code/internal/adapter/inbound/cityhdl"
	"github.com/gigihprasetyo/backend-standard-code/internal/adapter/outbound/cityrps"
	"github.com/gigihprasetyo/backend-standard-code/internal/core/services/citysvc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type GrpcServer struct {
	Postgres *gorm.DB
	Log      *zap.Logger
	R        *grpc.Server
}

func (h *GrpcServer) SetupRpc() {
	cityRep := cityrps.NewCityPostgres(h.Postgres)

	cityService := citysvc.NewCityService(h.Log, cityRep)

	cityhdl.NewCityGrpc(h.R, cityService)
}
