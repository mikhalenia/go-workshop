package order

import (
	"fmt"
	"logs"
	"payment"
	"sync"
)

type IPaid interface {
	Paid() error
}

type Order struct {
	id   int
	user int
	auto int
	sum  int
	//from   Time
	//to     Time
	status bool
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

func (o Order) Payment(pType string) error {
	var p IPaid
	switch pType {
	case ERIP:
		p = IPaid(payment.CreateEripOrder(o.id))
	case VISA:
		p = IPaid(payment.CreateVisaOrder(o.id))
	default: // Cash payment
		p = IPaid(o)
	}

	return p.Paid()
}

func (o Order) Paid() (err error) { // Cash payment
	//
	if err == nil {
		logs.Logs(fmt.Sprintf("Order %d is paid", o.id))
	}
	return
}

func (c *Catalog) SetOrder(order Order) error { // got some value for create Order
	//	order := Order{}
	//	order.id = c.createId()
	//	for i, val := range value {
	//		switch i {
	//		case 0:
	//			order.user = val.(int)
	//		case 1:
	//			order.auto = val.(int)
	//		case 2:
	//			order.sum = val.(int)
	//		}
	//	}
	logs.Logs(fmt.Sprintf("Created order: %v", order))
	c.Lock()
	c.Table[order.id] = order
	c.Unlock()
	return nil
}

func (c *Catalog) GetOrder(id int) Order {
	return c.Table[id]
}

func (c *Catalog) DeleteOrder(id int) {
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
