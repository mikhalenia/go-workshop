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

type Stack struct {
	sync.Mutex
	Table map[int]Order
}

func CreateStack() Stack {
	return Stack{
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

func (s *Stack) SetOrder(value ...interface{}) error { // got some value for create Order
	order := Order{}
	order.id = s.createId()
	for i, val := range value {
		switch i {
		case 0:
			order.user = val.(int)
		case 1:
			order.auto = val.(int)
		case 2:
			order.sum = val.(int)
		}
	}
	logs.Logs(fmt.Sprintf("Created order: %v", order))
	s.Lock()
	s.Table[order.id] = order
	s.Unlock()
	return nil
}

func (s *Stack) GetOrder(id int) Order {
	return s.Table[id]
}

func (s *Stack) DeleteOrder(id int) {
	s.Lock()
	delete(s.Table, id)
	s.Unlock()
}

func (s *Stack) createId() (id int) {
	i := 1
	for {
		if _, ok := s.Table[i]; !ok {
			id = i
			logs.Logs(fmt.Sprintf("Created id: %d", id))
			return
		}
		i++
	}
}
