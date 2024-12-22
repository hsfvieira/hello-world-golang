package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hsfvieira/hello-world-golang/usecases"
)

func main() {
	http.HandleFunc("GET /api/users/{username}", getUserByUsernameController)

	http.HandleFunc("GET /api/users", getUsersController)

	http.HandleFunc("POST /api/users", postUsersController)

	http.HandleFunc("GET /users", viewController("templates/users.html", &usecases.Users))

	http.HandleFunc("GET /users/new", viewController("templates/user-new.html", nil))

	http.HandleFunc("POST /users/new", postUsersFormController)

	http.HandleFunc("GET /users/{username}", getUserByUsernameFormController)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	http.ListenAndServe(port, nil)
}
