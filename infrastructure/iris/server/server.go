package server

import (
	"github.com/VulpesFerrilata/user/infrastructure/iris/router"
	"github.com/kataras/iris/v12"
)

func NewServer(router router.Router) *iris.Application {
	app := iris.Default()
	router.InitRoutes(app)
	return app
}
