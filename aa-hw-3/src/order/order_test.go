package order

import (
	"testing"
)

var orders Catalog

func TestOrder(t *testing.T) {
	var order Order
	orders = CreateCatalog()

	for i := 0; i < 10; i++ {
		order = Order{orders.createId(), i, i, i, false}
		orders.SetOrder(order)
	}

	o := orders.GetOrder(5)
	orders.DeleteOrder(o.id)

	order = Order{orders.createId(), 100, 100, 100, false}
	orders.SetOrder(order)

	o = orders.GetOrder(5)

	if o.id != 5 {
		t.Errorf("Error Test Number Order. Want = 5 and Got = '%d'.", o.id)
	}

	if len(orders.Table) != 10 {
		t.Errorf("Error Test Len Users. Want = 10 and Got = '%d'.", len(orders.Table))
	}

	order = Order{orders.createId(), 11, 11, 11, false}
	orders.SetOrder(order)
	o = orders.GetOrder(len(orders.Table))

	if o.id != 11 {
		t.Errorf("Error Test Number Order. Want = 11 and Got = '%d'.", o.id)
	}

	if len(orders.Table) != 11 {
		t.Errorf("Error Test Len Users. Want = 11 and Got = '%d'.", len(orders.Table))
	}
}
