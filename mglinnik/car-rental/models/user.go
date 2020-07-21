package models

import (
	"time"
)

// User ...
type User struct {
	ID       int64
	Name     string
	Password string
	Car      []*Car // rented car
}

func NewUser(name,password string) *User {
	user := new(User)
	user.ID = GenerateID()
	user.Name = name
	user.Password = password
	return user
}

func (u *User) RentCar(brand, model string) {
	rented := NewCar(brand, model)
	u.Car = append(u.Car, rented)
}

func (u *User) GetID() int64 {
	return u.ID
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetCar() []Car {
	return []Car{}
}

func GenerateID() int64 {
	uniqueNumber := time.Now().UnixNano() / (1 << 22)
	return uniqueNumber
}