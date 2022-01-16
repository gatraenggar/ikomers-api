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
}

type RegisterUserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type RegisterUser struct {
	UserRepository     model.UserRepository
	SecurityRepository helperModel.SecurityRepository
}

func NewRegisterUserUsecase(u model.UserRepository, s helperModel.SecurityRepository) *RegisterUser {
	return &RegisterUser{
		UserRepository:     u,
		SecurityRepository: s,
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
	}
	err = user.ValidateFields()
	if err != nil {
		return nil, err
	}

	userID, err := r.SecurityRepository.GenerateID(ctx)
	if err != nil {
		return nil, err
	}
	user.ID = userID

	hashedPass, err := r.SecurityRepository.HashPassword(ctx, user.Password)
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
	}, nil
}
