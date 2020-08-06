package jsons

import (
	"testing"
)

type JsonTest struct {
	Cars         []Car
	c1, c2, c3   Car
	firstResult  []string
	secondResult []string

}
func Test(t *testing.T) {
	toJsonTest := []JsonTest{
		{
			[]Car{
				{"1234AB-7", "audi", "black", 2018, 47, false},
			},
			Car{"1574AP-7", "fiat", "black", 2010, 17, false},
			Car{"6664AB-7", "reno", "blue", 2020, 50, true},
			Car{"3566AC-4", "bentley", "grey", 2020, 100, false},
			[]string{
				`{"number":"1234AB-7","brand":"audi","color":"black","year":2018,"cost":47,"available":false}`,
			},
			[]string{
				`{"number":"1234AB-7","brand":"audi","color":"black","year":2018,"cost":47,"available":false}`,
				`{"number":"1574AP-7","brand":"fiat","color":"black","year":2010,"cost":17,"available":false}`,
				`{"number":"6664AB-7","brand":"reno","color":"blue","year":2020,"cost":50,"available":true}`,
				`{"number":"3566AC-4","brand":"bentley","color":"grey","year":2020,"cost":100,"available":false}`,
			},
		},
		{
			[]Car{
				{"5445AP-3", "bmw", "white", 2011, 29, true},
				{"9484AK-5", "chevrolet", "green", 2019, 67, true},
			},
			Car{"1574AP-7", "fiat", "black", 2010, 17, false},
			Car{"6664AB-7", "reno", "blue", 2020, 50, true},
			Car{"1234AB-7", "audi", "black", 2018, 47, false},
			[]string{
				`{"number":"5445AP-3","brand":"bmw","color":"white","year":2011,"cost":29,"available":true}`,
				`{"number":"9484AK-5","brand":"chevrolet","color":"green","year":2019,"cost":67,"available":true}`,
			},
			[]string{
				`{"number":"5445AP-3","brand":"bmw","color":"white","year":2011,"cost":29,"available":true}`,
				`{"number":"9484AK-5","brand":"chevrolet","color":"green","year":2019,"cost":67,"available":true}`,
				`{"number":"1574AP-7","brand":"fiat","color":"black","year":2010,"cost":17,"available":false}`,
				`{"number":"6664AB-7","brand":"reno","color":"blue","year":2020,"cost":50,"available":true}`,
				`{"number":"1234AB-7","brand":"audi","color":"black","year":2018,"cost":47,"available":false}`,
			},
		},
	}
	for _, jt := range toJsonTest {
		var cars []Car
		for i, car := range jt.Cars {
			jsonString1 := toJson(car)
			cars = append(cars, fromJson(jsonString1))
			if jt.firstResult[i] != jsonString1 {
				t.Errorf("ERROR: Expected: %v, got: %v", jt.firstResult[i], jsonString1)
			}
		}
		cars = append(cars, jt.c1)
		cars = append(cars, jt.c2)
		cars = append(cars, jt.c3)
		for i, car := range cars {
			jsonString2 := toJson(car)
			if jt.secondResult[i] != jsonString2 {
				t.Errorf("ERROR: Expected: %v, got: %v", jt.secondResult[i], jsonString2)
			}
		}
	}
}