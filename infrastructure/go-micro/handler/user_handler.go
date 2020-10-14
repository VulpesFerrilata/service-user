package handler

import (
	"context"

	"github.com/VulpesFerrilata/grpc/protoc/user"
	"github.com/VulpesFerrilata/user/internal/usecase/interactor"
	"github.com/VulpesFerrilata/user/internal/usecase/request"
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
	userRequest := new(request.UserRequest)
	userRequest.ID = int(userRequestPb.GetID())

	userResponse, err := uh.userInteractor.GetUserById(ctx, userRequest)
	if err != nil {
		return err
	}

	userResponsePb.ID = int64(userResponse.ID)

	return nil
}

func (uh userHandler) GetUserByCredential(ctx context.Context, credentialRequestPb *user.CredentialRequest, userResponsePb *user.UserResponse) error {
	credentialRequest := new(request.CredentialRequest)
	credentialRequest.Username = credentialRequestPb.GetUsername()
	credentialRequest.Password = credentialRequestPb.GetPassword()

	userResponse, err := uh.userInteractor.GetUserByCredential(ctx, credentialRequest)
	if err != nil {
		return err
	}

	userResponsePb.ID = int64(userResponse.ID)

	return nil
}
