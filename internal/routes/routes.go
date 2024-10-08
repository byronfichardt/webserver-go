package routes

import (
	"database/sql"
	"net/http"
	users "webServer/internal/controllers"
)

func Router(db *sql.DB) {
	http.HandleFunc("POST /user/create", users.CreateUser(db))
	http.HandleFunc("POST /user/edit", users.EditUser(db))
	http.HandleFunc("GET /user", users.ListUsers(db))
}
