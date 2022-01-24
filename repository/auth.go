package repository

import (
	"context"
	"ikomers-be/database"
	"ikomers-be/model"
	"log"

	"gorm.io/gorm"
)

type mySqlAuthRepo struct {
	DB *gorm.DB
}

func NewMySqlAuthRepo(gormDB *gorm.DB) model.AuthRepository {
	db, err := database.NewDB()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &mySqlAuthRepo{
		DB: db,
	}
}

func (m *mySqlAuthRepo) AddRefreshToken(ctx context.Context, refreshToken string) error {
	res := m.DB.Create(database.Auth{RefreshToken: refreshToken})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
