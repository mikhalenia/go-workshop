package carrent

import (
	"studying_go_high_technologies_park/homework_3/cars"
	"studying_go_high_technologies_park/homework_3/renters"
)

type CarRentalSystem struct {
	Cars []*cars.Car
	Renters []*renters.Renter
}

func (agency CarRentalSystem) RegisterNewRenter(name, surname, phoneNumber, emailAddress string){
	renter := new(renters.Renter)
	renter.Name = name
	renter.Surname = surname
	renter.PhoneNumber = phoneNumber
	renter.EmailAddress = emailAddress
	agency.Renters = append(agency.Renters, renter)
}

func (agency CarRentalSystem) AddNewCar(number, brand, color string, year int, cost float32)  {
	car := new(cars.Car)
	car.Number = number
	car.Brand = brand
	car.Color = color
	car.Year = year
	car.Cost = cost
	agency.Cars = append(agency.Cars, car)
}
func (agency CarRentalSystem) RentCar(renter renters.Renter, rentStart, rentEnd string)  {
	// Looking for a car suitable for the interval [rentStart, rentEnd]
	// If have found several, choose with the lowest cost
	// if haven't found - error message
}





