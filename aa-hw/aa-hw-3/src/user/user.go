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

type Catalog struct {
	sync.Mutex
	Table map[int]User
}

func CreateCatalog() Catalog {
	return Catalog{
		Table: make(map[int]User),
	}
}

func (c *Catalog) SetUser(user User) error { // got some value for create Auto
	//	user := User{}
	//	user.id = c.createId()
	//	for i, val := range value {
	//		switch i {
	//		case 0:
	//			user.name = vac.(string)
	//		case 1:
	//			user.surname = vac.(string)
	//		case 2:
	//			user.login = vac.(string)
	//		case 3:
	//			user.password = vac.(string)
	//		}
	//	}
	logs.Logs(fmt.Sprintf("Created user: %v", user))
	c.Lock()
	c.Table[user.id] = user
	c.Unlock()
	return nil
}

func (c *Catalog) GetUser(id int) User {
	return c.Table[id]
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
