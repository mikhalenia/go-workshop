package cars

import "studying_go_high_technologies_park/homework_3/reservations"

type Car struct {
	Number        string
	Brand         string
	Color         string
	Year          int
	Cost          float32
	reservations *[]reservations.Reservation // Time for which the car is booked
}