package order

import (
	"testing"
)

var orders Stack

func TestOrder(t *testing.T) {

	orders = CreateStack()
	for i := 0; i < 10; i++ {
		orders.SetOrder(i, i, i)
	}

	o := orders.GetOrder(5)
	orders.DeleteOrder(o.id)

	orders.SetOrder(100, 100, 100)

	o = orders.GetOrder(5)

	if o.id != 5 {
		t.Errorf("Error Test Number Order. Want = 5 and Got = '%d'.", o.id)
	}

	if len(orders.Table) != 10 {
		t.Errorf("Error Test Len Users. Want = 10 and Got = '%d'.", len(orders.Table))
	}

	orders.SetOrder(11, 11, 11)
	o = orders.GetOrder(len(orders.Table))

	if o.id != 11 {
		t.Errorf("Error Test Number Order. Want = 11 and Got = '%d'.", o.id)
	}

	if len(orders.Table) != 11 {
		t.Errorf("Error Test Len Users. Want = 11 and Got = '%d'.", len(orders.Table))
	}
}
