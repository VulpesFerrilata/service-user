package handler

import (
	"context"

	"github.com/VulpesFerrilata/grpc/protoc/user"
	"github.com/VulpesFerrilata/user/infrastructure/go-micro/viewmodel"
	"github.com/VulpesFerrilata/user/internal/usecase/interactor"
)

func NewUserHandler(userInteractor interactor.UserInteractor) user.UserHandler {
	return userHandler{
		userInteractor: userInteractor,
	}
}

type userHandler struct {
	userInteractor interactor.UserInteractor
}

func (uh userHandler) GetUserById(ctx context.Context, userRequestPb *user.UserRequest, userResponsePb *user.UserResponse) error {
	userRequestVM := viewmodel.NewUserRequest(userRequestPb)

	userDTO, err := uh.userInteractor.GetUserById(ctx, userRequestVM.ToUserForm())
	if err != nil {
		return err
	}

	userResponseVM := viewmodel.NewUserResponse(userResponsePb)
	userResponseVM.FromUserDTO(userDTO)
	return nil
}

func (uh userHandler) GetUserByCredential(ctx context.Context, credentialRequestPb *user.CredentialRequest, userResponsePb *user.UserResponse) error {
	credentialRequestVM := viewmodel.NewCredentialRequest(credentialRequestPb)

	userDTO, err := uh.userInteractor.GetUserByCredential(ctx, credentialRequestVM.ToLoginForm())
	if err != nil {
		return err
	}

	userResponseVM := viewmodel.NewUserResponse(userResponsePb)
	userResponseVM.FromUserDTO(userDTO)
	return nil
}
