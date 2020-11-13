package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Sqlite driver
)

// Connect connect to the database.
func Connect() *sql.DB {
	var err error

	// Connect to database
	db, err := sql.Open("sqlite3", "./database/database.sqlite")
	if err != nil {
		panic(err)
	}

	// Test the database connexion
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// Migrate the table
	_, err = db.Exec("CREATE TABLE `user_pokemon` (`user_id` VARCHAR(55), `pokemon_id` VARCHAR(55));")
	if err != nil {
		fmt.Println("Table already exist")
	}

	return db
}
