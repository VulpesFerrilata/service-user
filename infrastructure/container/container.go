package container

import (
	"github.com/VulpesFerrilata/library/config"
	"github.com/VulpesFerrilata/library/pkg/database"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/VulpesFerrilata/library/pkg/translator"
	"github.com/VulpesFerrilata/library/pkg/validator"
	"github.com/VulpesFerrilata/user/infrastructure/go-micro/handler"
	"github.com/VulpesFerrilata/user/infrastructure/iris/controller"
	"github.com/VulpesFerrilata/user/infrastructure/iris/router"
	"github.com/VulpesFerrilata/user/infrastructure/iris/server"
	"github.com/VulpesFerrilata/user/internal/domain/repository"
	"github.com/VulpesFerrilata/user/internal/domain/service"
	"github.com/VulpesFerrilata/user/internal/usecase/interactor"
	"go.uber.org/dig"
)

func NewContainer() *dig.Container {
	container := dig.New()

	//--Config
	container.Provide(config.NewConfig)
	container.Provide(config.NewJwtConfig)

	//--Domain
	container.Provide(repository.NewUserRepository)
	container.Provide(service.NewUserService)
	//--Usecase
	container.Provide(interactor.NewUserInteractor)

	//--Utility
	container.Provide(database.NewGorm)
	container.Provide(translator.NewTranslator)
	container.Provide(validator.NewValidate)

	//--Middleware
	container.Provide(middleware.NewTransactionMiddleware)
	container.Provide(middleware.NewTranslatorMiddleware)
	container.Provide(middleware.NewErrorHandlerMiddleware)

	//--Controller
	container.Provide(controller.NewUserController)
	//--Router
	container.Provide(router.NewRouter)
	//--Server
	container.Provide(server.NewServer)

	//--Grpc
	container.Provide(handler.NewUserHandler)

	return container
}
