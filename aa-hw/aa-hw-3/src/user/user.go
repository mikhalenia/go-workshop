package user

import (
	"errors"
	"fmt"
	"logs"
	"sync"
)

type UserDB interface {
	Login(string, string) User
	Logout(int) bool
	DeleteUser(int)
	AddUser(User) error
}

type User struct {
	Id       int
	Name     string
	Surname  string
	Login    string
	Password string
	Email    string
	IP       string
	Status   string
}

type Catalog struct {
	sync.Mutex
	Table map[int]User
}

func CreateCatalog() Catalog {
	return Catalog{
		Table: make(map[int]User),
	}
}

func (c *Catalog) AddUser(user User) error { // got some value for create Auto
	logs.Logs(fmt.Sprintf("Created user: %v", user))
	c.Lock()
	c.Table[user.Id] = user
	c.Unlock()
	return nil
}

func (c *Catalog) Login(login string, password string) (*User, error) {
	//
	return nil, errors.New("Login failed")
}

func (c *Catalog) Logout(id int) bool {
	c.Lock()
	u := c.Table[id]
	u.Status = "offline"
	c.Unlock()
	return true
}

func (c *Catalog) DeleteUser(id int) {
	c.Lock()
	delete(c.Table, id)
	c.Unlock()
}

func (c *Catalog) createId() (id int) {
	i := 1
	for {
		if _, ok := c.Table[i]; !ok {
			id = i
			logs.Logs(fmt.Sprintf("Created id: %d", id))
			return
		}
		i++
	}
}
