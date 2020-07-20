package payment

type VisaOrder struct {
	id int
	// also some data
}

func CreateVisaOrder(val ...int) (order VisaOrder) {
	order = VisaOrder{val[0]}
	// add val to order
	return
}

func (o VisaOrder) Paid() error {
	//
	return nil
}
