package cmd

import (
	"fmt"
	"os"
)

func InitServer() {
	args := os.Args
	if len(args) != 2 {
		printCommands()
		return
	}

	switch args[1] {
	case "migrate":
		fmt.Println("Migrating the database")

	case "migrate_test":
		fmt.Println("Migrating the testing database...")
		MigrateTable()

	case "server":
		fmt.Println("Running the server...")

	default:
		fmt.Println("You seems to miss-typed the command...")
		printCommands()
	}
}

func printCommands() {
	fmt.Println("[Commands]")
	fmt.Println("go run main.go migrate \t\t : to migrate database")
	fmt.Println("go run main.go migrate_test \t : to migrate testing database")
	fmt.Println("go run main.go server \t\t : to run the server")
}
