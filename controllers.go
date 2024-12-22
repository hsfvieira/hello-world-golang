package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/hsfvieira/hello-world-golang/usecases"
)

func getUserByUsernameController(w http.ResponseWriter, r *http.Request) {
	userFiltered := usecases.FilterUserByUsername(r.PathValue("username"))

	w.Header().Add("Content-Type", "application/json")
	userJson, _ := json.Marshal(userFiltered)
	w.Write([]byte(userJson))
}

func postUsersController(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	var newUser usecases.User
	json.Unmarshal(data, &newUser)

	err := usecases.CreateNewUser(newUser)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(400)
		fmt.Fprintf(w, "{\"msg\": \"%s\"}", err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte("{\"msg\": \"User created\"}"))
}

func getUsersController(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	usersJson, _ := json.Marshal(usecases.GetAllUsers())
	w.Write([]byte(usersJson))
}

func viewController(templatePath string, data any) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles(templatePath)
		tmpl.Execute(w, data)
	}
}

func postUsersFormController(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 10)

	var newUser usecases.User

	newUser.Username = r.MultipartForm.Value["username"][0]
	newUser.Email = r.MultipartForm.Value["email"][0]
	newUser.Name = r.MultipartForm.Value["name"][0]

	err := usecases.CreateNewUser(newUser)
	if err != nil {
		tmpl, _ := template.ParseFiles("templates/user-new.html")
		tmpl.Execute(w, err.Error())
		return
	}

	tmpl, _ := template.ParseFiles("templates/user-new.html")
	tmpl.Execute(w, nil)
}

func getUserByUsernameFormController(w http.ResponseWriter, r *http.Request) {
	userFiltered := usecases.FilterUserByUsername(r.PathValue("username"))

	tmpl, _ := template.ParseFiles("templates/user-get.html")
	tmpl.Execute(w, userFiltered)
}
