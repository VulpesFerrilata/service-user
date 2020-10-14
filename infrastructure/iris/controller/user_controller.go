package controller

import (
	"github.com/VulpesFerrilata/user/internal/usecase/interactor"
	"github.com/VulpesFerrilata/user/internal/usecase/request"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController interface {
	PostRegister(ctx iris.Context) interface{}
}

func NewUserController(userInteractor interactor.UserInteractor) UserController {
	return &userController{
		userInteractor: userInteractor,
	}
}

type userController struct {
	userInteractor interactor.UserInteractor
}

func (uc userController) PostRegister(ctx iris.Context) interface{} {
	registerRequest := new(request.RegisterRequest)

	if err := ctx.ReadJSON(registerRequest); err != nil {
		return err
	}

	userResponse, err := uc.userInteractor.Register(ctx.Request().Context(), registerRequest)
	if err != nil {
		return err
	}

	return mvc.Response{
		Code:   iris.StatusCreated,
		Object: userResponse,
	}
}
