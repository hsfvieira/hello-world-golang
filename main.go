package main

import (
	"net/http"

	"github.com/hsfvieira/hello-world-golang/usecases"
)

func main() {
	http.HandleFunc("GET /api/users/{username}", getUserByUsernameController)

	http.HandleFunc("GET /api/users", getUsersController)

	http.HandleFunc("POST /api/users", postUsersController)

	http.HandleFunc("GET /users", viewController("templates/users.html", &usecases.Users))

	http.HandleFunc("GET /user/new", viewController("templates/user-new.html", nil))

	http.HandleFunc("POST /user/new", postUsersFormController)

	http.HandleFunc("GET /user/{username}", getUserByUsernameFormController)

	http.ListenAndServe(":8080", nil)
}
