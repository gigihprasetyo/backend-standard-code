package app

import (
	"github.com/gigihprasetyo/backend-standard-code/internal/adapter/inbound/userhdl"
	"github.com/gigihprasetyo/backend-standard-code/internal/adapter/outbound/authrps"
	"github.com/gigihprasetyo/backend-standard-code/internal/adapter/outbound/graphfacebookrps"
	"github.com/gigihprasetyo/backend-standard-code/internal/adapter/outbound/productrps"
	"github.com/gigihprasetyo/backend-standard-code/internal/adapter/outbound/profilerps"

	"github.com/gigihprasetyo/backend-standard-code/internal/adapter/outbound/userbankrps"
	"github.com/gigihprasetyo/backend-standard-code/internal/adapter/outbound/userrps"
	"github.com/gigihprasetyo/backend-standard-code/internal/core/services/usersvc"
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
	userRep := userrps.NewUserPostgres(h.Postgres)
	profileRep := profilerps.NewProfilePostgres(h.Postgres)
	myBankRep := userbankrps.NewUserBankPostgres(h.Postgres)
	productRep := productrps.NewProductPostgres(h.Postgres)
	authRepo := authrps.NewAuthGrpcRepository(h.Log)
	graphFacebookRepo := graphfacebookrps.NewGraphFacebookRepo()

	userService := usersvc.NewUserService(h.Log, userRep, profileRep, myBankRep, productRep, authRepo, graphFacebookRepo)

	userhdl.NewUserGrpc(h.R, userService)
}
