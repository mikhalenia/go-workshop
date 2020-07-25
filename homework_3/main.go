package main

import (
	"studying_go_high_technologies_park/homework_3/carrent"
)
func main(){
	carRentalAgency := new(carrent.CarRentalSystem)
	carRentalAgency.AddNewCar("1234AB-7", "audi", "чёрный", 2015, 47)
	carRentalAgency.AddNewCar("8682AP-7", "reno", "синий", 2018, 35)
	carRentalAgency.AddNewCar("23454AB-5", "mercedes", "белый", 2015, 80)
	carRentalAgency.AddNewCar("7777AB-7", "geely", "серебристый", 2019, 120)
	carRentalAgency.RegisterNewRenter("Иван", "Иванов", "+375337777777", "ivanov@gmail.com")
}
