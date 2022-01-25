package db_test

import (
	"ikomers-be/database"
	"ikomers-be/model"
	"log"

	"gorm.io/gorm"
)

type authTableTestHelper struct {
	DB *gorm.DB
}

func (t *authTableTestHelper) GetRefreshToken(refreshToken string) ([]model.Auth, error) {
	res := t.DB.First(&database.Auth{}, "refresh_token = ?", refreshToken)
	if res.Error != nil {
		return nil, res.Error
	}

	rows, err := res.Rows()
	if err != nil {
		log.Println("res.rows()")
		return nil, err
	}
	defer rows.Close()

	rowsResult := make([]model.Auth, 0)
	for rows.Next() {
		var auth model.Auth

		if err := rows.Scan(&auth.RefreshToken); err != nil {
			return nil, err
		}
		rowsResult = append(rowsResult, model.Auth{
			RefreshToken: auth.RefreshToken,
		})
	}
	if err := rows.Err(); err != nil {
		log.Println("rows.Err()")
		return nil, err
	}

	return rowsResult, err
}

func (t *authTableTestHelper) CleanTable() {
	t.DB.Exec("DELETE FROM auths WHERE 1=1")
}

func NewAuthTableTestHelper() *authTableTestHelper {
	db, err := database.NewDB()
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &authTableTestHelper{DB: db}
}
