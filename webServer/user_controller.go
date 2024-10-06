package main

import (
	"net/http"
	"time"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

type User struct {
	ID        int       `json:"id"`
    Username  string    `json:"username"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)

		if err != nil {
			log.Fatal(err)
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
        w.Write([]byte("User created successfully"))
	}
}

func ListUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := "SELECT id, username, created_at, updated_at FROM users"
        rows,err := db.Query(query)

		if err != nil {
			http.Error(w, "Failed to query users", http.StatusInternalServerError)
            return
		}

		defer rows.Close()

		var users []User

		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt); err != nil {
				http.Error(w, "Failed to scan user", http.StatusInternalServerError)
                return
			} 
			users = append(users, user)
		}

		if err := rows.Err(); err != nil {
            http.Error(w, "Error iterating over users", http.StatusInternalServerError)
            return
        }

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(users); err != nil {
			http.Error(w, "Failed to encode user to JSON", http.StatusInternalServerError)
		}
	}
}

func EditUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		_, err := db.Exec("UPDATE users SET password = ? WHERE username = ?", password, username)
        if err != nil {
            http.Error(w, "Failed to update user", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("User updated successfully"))
	}
}