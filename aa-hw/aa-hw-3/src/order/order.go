package order

import (
	//"auto"
	"fmt"
	"logs"
	"payment"
	"sync"
	//"user"
)

type LockAuto interface {
	Lock() bool
	Unlock()
}

/*func (c *Auto) Lock() bool {
	//
}

func (c *Auto) Unlock() {

}*/

/*type OrderDB interface {
	CreateOrder(User, Auto) Order
	AddOrder(Order)
}*/

type IPaid interface {
	Paid() error
}

type Order struct {
	Id     int
	UserId int
	AutoId int
	Sum    int
	Status string
}

const ERIP = "erip"
const VISA = "visa"

type Catalog struct {
	sync.Mutex
	Table map[int]Order
}

func CreateCatalog() Catalog {
	return Catalog{
		Table: make(map[int]Order),
	}
}

func (o *Order) Payment(pType string) error {
	var p IPaid
	switch pType {
	case ERIP:
		p = IPaid(payment.CreateEripOrder(o.Id))
	case VISA:
		p = IPaid(payment.CreateVisaOrder(o.Id))
	default: // Cash payment
		p = IPaid(o)
	}

	return p.Paid()
}

func (o *Order) Paid() (err error) { // Cash payment
	//
	if err == nil {
		logs.Logs(fmt.Sprintf("Order %d is paid", o.Id))
	}
	return
}

func (c *Catalog) AddOrder(order Order) error { // got some value for create Order
	logs.Logs(fmt.Sprintf("Created order: %v", order))
	c.Lock()
	c.Table[order.Id] = order
	c.Unlock()
	return nil
}

/*func (c *Catalog) CreateOrder(user User, auto Auto) Order {
	//
	return nil
}*/

func (c *Catalog) CreateId() (id int) {
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
