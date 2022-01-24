package database

import "gorm.io/gorm"

func MigrateTable(m gorm.Migrator) {
	m.CreateTable(Auth{})
	m.CreateTable(User{})
}
