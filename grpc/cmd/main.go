package main

import (
	"log"
	"net"

	"cms/grpc/helpers/database"
	"cms/grpc/helpers/xvalidator"
	_middelware "cms/grpc/module/auth/delivery/middleware_grpc"

	"cms/config"
	slogLogger "cms/grpc/helpers/logger"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_authGrpc "cms/grpc/module/auth/delivery/grpc"

	"google.golang.org/grpc"
)

func main() {
	validate, _ := xvalidator.NewValidator()
	conf := config.InitConfig(validate)
	slogLogger.SetupLogger(&slogLogger.Config{
		CurrentEnv: conf.AppEnv.CurrentEnv,
		LogPath:    conf.AppEnv.LogFilePath,
	})
	env := conf.AppEnv.CurrentEnv
	if env == "" {
		log.Println("Service RUN on DEBUG mode")

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Use(middleware.Recover())
	dbGorm := database.NewDatabase("postgres", &database.Config{
		DbHost: conf.Database.Pghost,
		DbUser: conf.Database.Pguser,
		DbPass: conf.Database.Pgpassword,
		DbName: conf.Database.Pgdatabase,
		DbPort: "5432",
	})

	repositoryAuth := _authRepository.NewAuthRepository()
	useCaseAuth := _authUseCase.NewAuthUseCase(dbGorm.GetDB(), repositoryAuth)

	middlewareJWT := _middelware.NewAuthenticationJWT(usecaseAuth, map[string]string{
		"AuthService": "LoginUser",
	})
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(middlewareJWT.JwtInterceptor),
	)

	_authGrpc.NewAuthService(grpcServer, usecaseAuth)

	httpPort := conf.AppEnv.HttpPort

	if httpPort == "" {
		httpPort = "9090"
	}

	lis, err := net.Listen("tcp", ":"+conf.AppEnv.HttpPort)
	if err != nil {
		log.Fatalf("failed to initialise the store: %s", err)
	}

	log.Fatal(grpcServer.Serve(lis))

}
