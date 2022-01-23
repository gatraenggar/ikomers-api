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
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Type      int    `json:"type"`
}

type RegisterUser struct {
	UserRepository  model.UserRepository
	SecurityManager helperModel.SecurityManager
}

func NewRegisterUserUsecase(u model.UserRepository, s helperModel.SecurityManager) *RegisterUser {
	return &RegisterUser{
		UserRepository:  u,
		SecurityManager: s,
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

	res, err := r.UserRepository.RegisterUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &RegisterUserResponse{
		ID:        res.ID,
		Email:     res.Email,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Type:      int(res.Type),
	}, nil
}
