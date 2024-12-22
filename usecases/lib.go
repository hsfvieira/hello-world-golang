package usecases

import "errors"

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var Users []User

func FilterUserByUsername(username string) User {
	var userFiltered User

	for _, user := range Users {
		if user.Username == username {
			userFiltered = user
		}
	}

	return userFiltered
}

func CreateNewUser(newUser User) error {
	userFiltered := FilterUserByUsername(newUser.Username)

	if userFiltered.Username != "" {
		return errors.New("username exists")
	}

	Users = append(Users, newUser)

	return nil
}

func GetAllUsers() []User {
	return Users
}
