package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// get the array of Users
func GetUsers() []*User {
	return users
}

//add the new User
func AddUser(u User) (User, error) {

	if u.ID != 0 {
		return User{}, errors.New("New User must not include id ")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int) (User, error) {

	for _, u := range users {
		if u.ID == id {
			return *u, nil

		}
	}

	return User{}, fmt.Errorf("User with Id '%v' not found ", id)
}

func UpdateUser(user User) (User, error) {

	for i, u := range users {
		if u.ID == user.ID {
			users[i] = &user
			return user, nil
		}

	}

	return user, fmt.Errorf("User of '%v id is not found", user.ID)
}

func RemoveUserById(id int) error {

	for i, u := range users {
		if u.ID == id {
			users = append(users[:1], users[i+1:]...)
			return nil
		}

	}

	return fmt.Errorf("user with '%v' id not found", id)
}
