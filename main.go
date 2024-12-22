package main

import (
	"net/http"
)

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func main() {
	http.HandleFunc("GET /api/users/{username}", getUserByUsername)

	http.HandleFunc("GET /api/users", getUsers)

	http.HandleFunc("POST /api/users", postUsers)

	http.ListenAndServe(":8080", nil)
}
