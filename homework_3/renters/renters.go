package renters

import (
	"studying_go_high_technologies_park/homework_3/cars"
)

type Renter struct {
	Name         string
	Surname      string
	PhoneNumber  string
	EmailAddress string
	RentedCar *cars.Car // Car that is rented now
	RentalHistory []string // Numbers of rented cars
}

