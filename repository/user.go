package repository

import (
	"context"
	"errors"
	"ikomers-be/model"

	"gorm.io/gorm"
)

type mySqlUserRepo struct {
	DB *gorm.DB
}

func NewMySqlUserRepo(gormDB *gorm.DB) model.UserRepository {
	return &mySqlUserRepo{
		DB: gormDB,
	}
}

func (m *mySqlUserRepo) GenerateID(ctx context.Context) (string, error) { return "", nil }

func (m *mySqlUserRepo) HashPassword(ctx context.Context, password string) (string, error) {
	return "", nil
}

func (m *mySqlUserRepo) CheckEmailAvailability(ctx context.Context, email string) error {
	var res model.User

	m.DB.Select("email").Where(&model.User{Email: email}).Find(&res)
	if res.Email == email {
		return errors.New("email is not available")
	}

	return nil
}

func (m *mySqlUserRepo) RegisterUser(ctx context.Context, user *model.User) (*model.User, error) {
	res := m.DB.Create(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}
