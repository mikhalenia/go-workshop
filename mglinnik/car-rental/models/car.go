package models

// Car ...
type Car struct {
	ID int64
	Brand string
	Model string
	Status bool
}

func NewCar(brand, model string) *Car {
	car := new(Car)
	car.ID = GenerateID()
	car.Brand = brand
	car.Model = model
	car.Status = true
	return car
}

func (c *Car) GetID() int64 {
	return c.ID
}

func (c*Car) GetBrand() string {
	return c.Brand
}

func (c *Car) GetModel() string {
	return c.Model
}

func (c *Car) IsAvailable() bool {
	return c.Status
}