package migrations

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

func Migrate() {
	db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/gotest?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
    );`

	// Executes the SQL query in our database. Check err to ensure there was no error.
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migrations applied successfully")
}
