package main

import (
	"go-workshop/mglinnik/car-rental/models"
)


func main() {
		NewCars := models.NewCar("Toyota", "Auris")
		NewUsers := models.NewUser("John", "qwerty123")

		var Data = new(struct {
			Users []*models.User
			Cars  []*models.Car
		})

		Data.Users = append(Data.Users, NewUsers)
		Data.Cars = append(Data.Cars, NewCars)
	}

