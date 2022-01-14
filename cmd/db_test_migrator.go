package cmd

import (
	"fmt"
	"ikomers-be/database"
	"ikomers-be/migration"
)

func MigrateTable() {
	db, err := database.NewDB("ikomers_test")
	if err != nil {
		panic(err)
	}

	migration.MigrateUserTable(db.Migrator())

	fmt.Println("Migrate success!")
}
