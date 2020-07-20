package main

import (
	"go-workshop/frogstair/twitter/users"
)

func main() {
	database := new(struct {
		Users []*users.User
	})
	usr := users.New("username", "password")
	database.Users = append(database.Users, usr)

	usr.Post("I like pineapple pizza")
}
