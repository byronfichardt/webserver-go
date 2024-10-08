package main

import (
	"database/sql"
	"log"
	"net/http"
	"webServer/internal/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := connectDb()

	routes.Router(db)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}

func connectDb() *sql.DB {
	db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/gotest?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
