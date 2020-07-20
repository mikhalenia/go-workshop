package payment

type EripOrder struct {
	id int
	// also some data
}

func CreateEripOrder(val ...int) (order EripOrder) {
	order = EripOrder{}
	for _, v := range val {
		order.id = v
	}
	return
}

func (o EripOrder) Paid() error {
	//
	return nil
}
