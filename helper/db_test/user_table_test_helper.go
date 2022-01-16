package db_test

import (
	"ikomers-be/database"
	"ikomers-be/model"
	"log"

	"gorm.io/gorm"
)

type userTableTestHelper struct {
	DB *gorm.DB
}

func (t *userTableTestHelper) AddUser(u *model.User) {
	res := t.DB.Create(u)
	if res.Error != nil {
		log.Fatalf(res.Error.Error())
	}
}

func (t *userTableTestHelper) GetUserByID(u *model.User, userID string) ([]model.User, error) {
	res := t.DB.First(u, "id = ?", userID)
	if res.Error != nil {
		return nil, res.Error
	}

	rows, err := res.Rows()
	if err != nil {
		log.Println("res.rows()")
		return nil, err
	}
	defer rows.Close()

	rowsResult := make([]model.User, 0)
	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
		); err != nil {
			return nil, err
		}
		rowsResult = append(rowsResult, model.User{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
	}
	if err := rows.Err(); err != nil {
		log.Println("rows.Err()")
		return nil, err
	}

	return rowsResult, err
}

func (t *userTableTestHelper) CleanTable() {
	t.DB.Exec("DELETE FROM users WHERE 1=1")
}

func NewUserTableTestHelper() *userTableTestHelper {
	userTable, err := database.NewDB()
	if err != nil {
		log.Fatalf(err.Error())
	}

	pool := &userTableTestHelper{DB: userTable}

	return pool
}
