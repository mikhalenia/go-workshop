package auto

import (
	"fmt"
	"testing"
)

var autos Catalog

//	id     int
//	mark   string
//	model  string
//	url    string
//	count  int
//	status bool

func TestAuto(t *testing.T) {
	var auto Auto
	autos = CreateCatalog()

	for i := 0; i < 10; i++ {
		auto = Auto{autos.createId(), fmt.Sprintf("mark %d", i), fmt.Sprintf("model %d", i), fmt.Sprintf("url %d", i), i, false}
		autos.SetAuto(auto)
	}

	u := autos.GetAuto(5)
	autos.DeleteAuto(u.id)

	auto = Auto{autos.createId(), fmt.Sprintf("mark %d", 0), fmt.Sprintf("model %d", 0), fmt.Sprintf("url %d", 0), 0, false}
	autos.SetAuto(auto)

	a := autos.GetAuto(5)

	if a.id != 5 {
		t.Errorf("Error Test Number Auto. Want = 5 and Got = '%d'.", a.id)
	}

	if len(autos.Table) != 10 {
		t.Errorf("Error Test Len Autos. Want = 10 and Got = '%d'.", len(autos.Table))
	}
	auto = Auto{autos.createId(), fmt.Sprintf("mark %d", 11), fmt.Sprintf("model %d", 11), fmt.Sprintf("url %d", 11), 11, false}
	autos.SetAuto(auto)
	a = autos.GetAuto(len(autos.Table))

	if a.id != 11 {
		t.Errorf("Error Test Number Auto. Want = 11 and Got = '%d'.", a.id)
	}

	if len(autos.Table) != 11 {
		t.Errorf("Error Test Len Autos. Want = 11 and Got = '%d'.", len(autos.Table))
	}

}
