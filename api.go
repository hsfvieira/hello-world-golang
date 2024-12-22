package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func getUserByUsername(w http.ResponseWriter, r *http.Request) {
	var userFiltered User

	for _, user := range users {
		if user.Username == r.PathValue("username") {
			userFiltered = user
		}
	}

	w.Header().Add("Content-Type", "application/json")
	userJson, _ := json.Marshal(userFiltered)
	w.Write([]byte(userJson))
}

func postUsers(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	var newUser User
	json.Unmarshal(data, &newUser)

	var userFiltered User

	for _, user := range users {
		if user.Username == newUser.Username {
			userFiltered = user
		}
	}

	if userFiltered.Username != "" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(400)
		w.Write([]byte("{\"msg\": \"Username exists\"}"))
		return
	}

	users = append(users, newUser)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write([]byte("{\"msg\": \"User created\"}"))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	usersJson, _ := json.Marshal(users)
	w.Write([]byte(usersJson))
}
