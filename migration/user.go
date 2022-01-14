package migration

import "gorm.io/gorm"

type User struct {
	ID        string  `gorm:"primaryKey"`
	Email     *string `gorm:"unique"`
	FirstName string  `gorm:"size:15"`
	LastName  string  `gorm:"size:15"`
	Password  string
}

func MigrateUserTable(m gorm.Migrator) {
	m.CreateTable(User{})
}
