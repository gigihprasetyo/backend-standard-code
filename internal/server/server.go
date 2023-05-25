package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"

	baseApp "github.com/gigihprasetyo/backend-standard-code/internal/app"

	"github.com/gigihprasetyo/backend-standard-code/pkg/logger"
	"github.com/gigihprasetyo/backend-standard-code/pkg/viper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	viperPkg "github.com/spf13/viper"
)

func Run() {

	//load config
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "../..")
	config := &viper.EnvConfig{
		FileName: "config",
		FileType: "yaml",
		Path:     basepath,
	}
	if err := config.ReadConfig(); err != nil {
		log.Fatal(err)
	}

	// //load connection config
	// pg, err := postgres.Connect()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// sqlDB, err := pg.DB()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer func(sqlDB *sql.DB) {
	// 	err := sqlDB.Close()
	// 	if err != nil {
	// 	}
	// }(sqlDB)

	//load connection redis
	// red, err := redis.Connect()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//initial sentry hook zap log
	// if viperPkg.GetString("server.mode") != "local" {
	// 	err := logger.SentryInit()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// fmt.Println("sentry pass")
	// zap, err := logger.Initialize()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//load fiber
	app := fiber.New(fiber.Config{
		IdleTimeout: 5,
	})
	app.Use(
		recover.New(),
		compress.New(),
		etag.New(),
		fiberlog.New(),
		cors.New(),
	)
	rh := &baseApp.Handlers{
		//Postgres: pg,
		R: app,
		//Logger:   zap,
		//Redis:    red,
	}
	rh.SetupRouter()

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(":" + viperPkg.GetString("server.port")); err != nil {
			log.Panicf("failed listen into port %v", err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	var _ = <-c // This blocks the main thread until an interrupt is received
	log.Println("gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	//sqlDB.Close()
	fmt.Println("services was successful shutdown.")
}

func RpcRun() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "../..")
	config := &viper.EnvConfig{
		FileName: "config",
		FileType: "yaml",
		Path:     basepath,
	}
	if err := config.ReadConfig(); err != nil {
		log.Fatal(err)
	}

	//load connection postgre
	// pg, err := postgres.Connect()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// sqlDB, err := pg.DB()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer func(sqlDB *sql.DB) {
	// 	err := sqlDB.Close()
	// 	if err != nil {
	// 	}
	// }(sqlDB)

	//logger setup
	// if viperPkg.GetString("server.mode") != "local" {
	// 	err := logger.SentryInit()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	zap, err := logger.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	lis, err := net.Listen("tcp", ":"+viperPkg.GetString("rpc.port"))
	if err != nil {
		log.Fatal(err)

	}
	srv := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zap),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zap),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	rh := &baseApp.GrpcServer{
		//Postgres: pg,
		Log: zap,
		R:   srv,
	}
	rh.SetupRpc()
	reflection.Register(srv)

	log.Println("rpc run in port " + viperPkg.GetString("rpc.port"))
	err = srv.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("rpc run in port " + viperPkg.GetString("rpc.port"))
}
