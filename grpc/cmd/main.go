package main

import (
	"log"
	"net"

	"cms/grpc/helpers/database"
	slogLogger "cms/grpc/helpers/logger"
	"cms/grpc/helpers/utils/converter"
	"cms/grpc/helpers/xvalidator"
	_middelware "cms/grpc/module/auth/delivery/middleware_grpc"

	"cms/grpc/config"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"google.golang.org/grpc"

	_userGrpc "cms/grpc/module/user/delivery/grpc"
	_userRepository "cms/grpc/module/user/repository/postgres"
	_userUseCase "cms/grpc/module/user/usecase"

	_articleGrpc "cms/grpc/module/article/delivery/grpc"
	_articleRepository "cms/grpc/module/article/repository/postgres"
	_articleUseCase "cms/grpc/module/article/usecase"

	_authGrpc "cms/grpc/module/auth/delivery/grpc"
	_authRepository "cms/grpc/module/auth/repository/postgres"
	_authUseCase "cms/grpc/module/auth/usecase"
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
		DbPort: converter.ToString(conf.Database.Pgport),
	})

	userRepository := _userRepository.NewUserRepository()
	userUseCase := _userUseCase.NewUserUseCase(dbGorm.GetDB(), userRepository)

	articleRepository := _articleRepository.NewArticleRepository()
	articleUseCase := _articleUseCase.NewArticleUseCase(dbGorm.GetDB(), articleRepository)

	authRepository := _authRepository.NewAuthRepository()
	authUseCase := _authUseCase.NewAuthUseCase(dbGorm.GetDB(), authRepository, userRepository)

	middlewareJWT := _middelware.NewAuthenticationJWT(authUseCase, map[string][]string{
		"AuthService": {"LoginUser", "RegisterUser"}, // ← ini yang benar
	})
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(middlewareJWT.JwtInterceptor),
	)

	_userGrpc.NewUserService(grpcServer, userUseCase)
	_articleGrpc.NewArticleService(grpcServer, articleUseCase)
	_authGrpc.NewAuthService(grpcServer, authUseCase)

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
