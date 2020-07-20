package user

import (
	"fmt"
	"testing"
)

var users Stack

func TestUser(t *testing.T) {

	users = CreateStack()
	for i := 0; i < 10; i++ {
		users.SetUser(fmt.Sprintf("name %d", i), fmt.Sprintf("surname %d", i), fmt.Sprintf("login %d", i), fmt.Sprintf("password %d", i))
	}

	u := users.GetUser(5)
	users.DeleteUser(u.id)

	users.SetUser(fmt.Sprintf("name %d", 0), fmt.Sprintf("surname %d", 0), fmt.Sprintf("login %d", 0), fmt.Sprintf("password %d", 0))

	u = users.GetUser(5)

	if u.id != 5 {
		t.Errorf("Error Test Number User. Want = 5 and Got = '%d'.", u.id)
	}

	if len(users.Table) != 10 {
		t.Errorf("Error Test Len Users. Want = 10 and Got = '%d'.", len(users.Table))
	}

	users.SetUser(fmt.Sprintf("name %d", 11), fmt.Sprintf("surname %d", 11), fmt.Sprintf("login %d", 11), fmt.Sprintf("password %d", 11))
	u = users.GetUser(len(users.Table))

	if u.id != 11 {
		t.Errorf("Error Test Number Order. Want = 11 and Got = '%d'.", u.id)
	}

	if len(users.Table) != 11 {
		t.Errorf("Error Test Len Users. Want = 11 and Got = '%d'.", len(users.Table))
	}

}
