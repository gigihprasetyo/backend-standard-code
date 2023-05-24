package app

import (
	"github.com/pasarin-tech/pasarin-core/internal/adapter/inbound/userhdl"
	"github.com/pasarin-tech/pasarin-core/internal/adapter/outbound/authrps"
	"github.com/pasarin-tech/pasarin-core/internal/adapter/outbound/graphfacebookrps"
	"github.com/pasarin-tech/pasarin-core/internal/adapter/outbound/productrps"
	"github.com/pasarin-tech/pasarin-core/internal/adapter/outbound/profilerps"

	"github.com/pasarin-tech/pasarin-core/internal/adapter/outbound/userbankrps"
	"github.com/pasarin-tech/pasarin-core/internal/adapter/outbound/userrps"
	"github.com/pasarin-tech/pasarin-core/internal/core/services/usersvc"
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
