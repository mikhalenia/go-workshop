package user

import (
	"fmt"
	"logs"
	"sync"
)

type User struct {
	id       int
	name     string
	surname  string
	login    string
	password string
	email    string
}

type Stack struct {
	sync.Mutex
	Table map[int]User
}

func CreateStack() Stack {
	return Stack{
		Table: make(map[int]User),
	}
}

func (s *Stack) SetUser(value ...interface{}) error { // got some value for create Auto
	user := User{}
	user.id = s.createId()
	for i, val := range value {
		switch i {
		case 0:
			user.name = val.(string)
		case 1:
			user.surname = val.(string)
		case 2:
			user.login = val.(string)
		case 3:
			user.password = val.(string)
		}
	}
	logs.Logs(fmt.Sprintf("Created user: %v", user))
	s.Lock()
	s.Table[user.id] = user
	s.Unlock()
	return nil
}

func (s *Stack) GetUser(id int) User {
	return s.Table[id]
}

func (s *Stack) DeleteUser(id int) {
	s.Lock()
	delete(s.Table, id)
	s.Unlock()
}

func (s *Stack) createId() (id int) {
	i := 1
	for {
		if _, ok := s.Table[i]; !ok {
			id = i
			logs.Logs(fmt.Sprintf("Created id: %d", id))
			return
		}
		i++
	}
}
