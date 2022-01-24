package cmd

import (
	"fmt"
	"ikomers-be/database"
)

func MigrateTable() {
	db, err := database.NewDB()
	if err != nil {
		panic(err)
	}

	database.MigrateTable(db.Migrator())

	fmt.Println("Migrate success!")
}
