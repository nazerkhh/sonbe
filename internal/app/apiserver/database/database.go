// internal/app/apiserver/database/database.go
package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DB is the global variable for the database connection
var DB *sql.DB

// InitDB initializes the database connection
func InitDB() error {
	connStr := "user=your_user dbname=your_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected to the database")

	DB = db
	return nil
}
