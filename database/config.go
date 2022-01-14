package database

import (
	"fmt"
	"ikomers-be/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(dbName string) (*gorm.DB, error) {
	conf := config.GetConfig()

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DbTestUser,
		conf.DbTestPass,
		conf.DbTestHost,
		conf.DbTestPort,
		conf.DbTestName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
