package main

import (
	"database/sql"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/server/grpc"

	"github.com/VulpesFerrilata/grpc/protoc/user"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/VulpesFerrilata/user/infrastructure/container"
	"github.com/micro/go-micro/v2"
)

func main() {
	container := container.NewContainer()

	if err := container.Invoke(func(userHandler user.UserHandler,
		transactionMiddleware *middleware.TransactionMiddleware,
		translatorMiddleware *middleware.TranslatorMiddleware,
		errorHandlerMiddleware *middleware.ErrorHandlerMiddleware) error {
		// New Service
		service := micro.NewService(
			micro.Server(
				grpc.NewServer(
					server.WrapHandler(errorHandlerMiddleware.HandlerWrapper),
					server.WrapHandler(translatorMiddleware.HandlerWrapper),
					server.WrapHandler(transactionMiddleware.HandlerWrapperWithTxOptions(&sql.TxOptions{})),
				),
			),
			micro.Name("boardgame.user.svc"),
			micro.Version("latest"),
		)

		// Initialise service
		service.Init()

		// Register Handler
		if err := user.RegisterUserHandler(service.Server(), userHandler); err != nil {
			return err
		}

		// Run service
		return service.Run()
	}); err != nil {
		log.Fatal(err)
	}
}
