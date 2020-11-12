package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Sqlite driver
)

var (
	// DB var
	DB *sql.DB
)

// Connect connect to the database.
func Connect() {
	var err error

	// Connect to database
	DB, err := sql.Open("sqlite3", "./database/database.sqlite")
	if err != nil {
		panic(err)
	}

	// Test the database connexion
	if err := DB.Ping(); err != nil {
		panic(err)
	}

	// Migrate the table
	_, err = DB.Exec("CREATE TABLE `user_pokemon` (`user_id` INTEGER, `pokemon_id` INTEGER);")
	if err != nil {
		fmt.Println("Table already exist")
	}
}
