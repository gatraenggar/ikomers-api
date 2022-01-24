package use_case

import (
	"context"
	"ikomers-be/model"
	helperModel "ikomers-be/model/helper"
)

type RegisterUserRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Type      int    `json:"type"`
}

type RegisterUserResponse struct {
	ID        string     `json:"id"`
	Email     string     `json:"email"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Type      int        `json:"type"`
	AuthToken model.Auth `json:"auth_token"`
}

type RegisterUser struct {
	AuthRepository  model.AuthRepository
	UserRepository  model.UserRepository
	SecurityManager helperModel.SecurityManager
	TokenManager    helperModel.TokenManager
}

func NewRegisterUserUsecase(
	a model.AuthRepository,
	u model.UserRepository,
	s helperModel.SecurityManager,
	t helperModel.TokenManager,
) *RegisterUser {
	return &RegisterUser{
		AuthRepository:  a,
		UserRepository:  u,
		SecurityManager: s,
		TokenManager:    t,
	}
}

func (r *RegisterUser) Execute(ctx context.Context, req *RegisterUserRequest) (*RegisterUserResponse, error) {
	err := r.UserRepository.CheckEmailAvailability(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
		Type:      model.UserType(req.Type),
	}
	err = user.ValidateFields()
	if err != nil {
		return nil, err
	}

	userID, err := r.SecurityManager.GenerateID(ctx)
	if err != nil {
		return nil, err
	}
	user.ID = userID

	hashedPass, err := r.SecurityManager.HashPassword(ctx, user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPass

	registeredUser, err := r.UserRepository.RegisterUser(ctx, user)
	if err != nil {
		return nil, err
	}

	accessToken, err := r.TokenManager.GenerateAccessToken(ctx, *user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := r.TokenManager.GenerateRefreshToken(ctx, *user)
	if err != nil {
		return nil, err
	}

	r.AuthRepository.AddRefreshToken(ctx, refreshToken)

	return &RegisterUserResponse{
		ID:        registeredUser.ID,
		Email:     registeredUser.Email,
		FirstName: registeredUser.FirstName,
		LastName:  registeredUser.LastName,
		Type:      int(registeredUser.Type),
		AuthToken: model.Auth{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
