package jsons

import (
	"testing"
)

type JsonTest struct {
	Cars         []Car
	c1, c2, c3   Car
	firstResult  string
	secondResult string

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
			"[\n{\n\"number\": \"1234AB-7\",\n\"brand\": \"audi\",\n\"color\": \"black\",\n\"year\": 2018,\n\"cost\": 47,\n\"available\": false\n}\n]",
			"[\n{\n\"number\": \"1234AB-7\",\n\"brand\": \"audi\",\n\"color\": \"black\",\n\"year\": 2018,\n\"cost\": 47,\n\"available\": false\n},\n{\n\"number\": \"1574AP-7\",\n\"brand\": \"fiat\",\n\"color\": \"black\",\n\"year\": 2010,\n\"cost\": 17,\n\"available\": false\n},\n{\n\"number\": \"6664AB-7\",\n\"brand\": \"reno\",\n\"color\": \"blue\",\n\"year\": 2020,\n\"cost\": 50,\n\"available\": true\n},\n{\n\"number\": \"3566AC-4\",\n\"brand\": \"bentley\",\n\"color\": \"grey\",\n\"year\": 2020,\n\"cost\": 100,\n\"available\": false\n}\n]",
		},
		{
			[]Car{
				{"5445AP-3", "bmw", "white", 2011, 29, true},
				{"9484AK-5", "chevrolet", "green", 2019, 67, true},
			},
			Car{"5445AP-3", "bmw", "white", 2011, 29, true},
			Car{"9484AK-5", "chevrolet", "green", 2019, 67, true},
			Car{"1234AB-7", "audi", "black", 2018, 47, false},
			"[\n{\n\"number\": \"5445AP-3\",\n\"brand\": \"bmw\",\n\"color\": \"white\",\n\"year\": 2011,\n\"cost\": 29,\n\"available\": true\n},\n{\n\"number\": \"9484AK-5\",\n\"brand\": \"chevrolet\",\n\"color\": \"green\",\n\"year\": 2019,\n\"cost\": 67,\n\"available\": true\n}\n]",
			"[\n{\n\"number\": \"5445AP-3\",\n\"brand\": \"bmw\",\n\"color\": \"white\",\n\"year\": 2011,\n\"cost\": 29,\n\"available\": true\n},\n{\n\"number\": \"9484AK-5\",\n\"brand\": \"chevrolet\",\n\"color\": \"green\",\n\"year\": 2019,\n\"cost\": 67,\n\"available\": true\n},\n{\n\"number\": \"5445AP-3\",\n\"brand\": \"bmw\",\n\"color\": \"white\",\n\"year\": 2011,\n\"cost\": 29,\n\"available\": true\n},\n{\n\"number\": \"9484AK-5\",\n\"brand\": \"chevrolet\",\n\"color\": \"green\",\n\"year\": 2019,\n\"cost\": 67,\n\"available\": true\n},\n{\n\"number\": \"1234AB-7\",\n\"brand\": \"audi\",\n\"color\": \"black\",\n\"year\": 2018,\n\"cost\": 47,\n\"available\": false\n}\n]",
		},
	}

	for _, jt := range toJsonTest {
		jsonString := toJson(jt.Cars)
		if jt.firstResult != jsonString {
			t.Errorf("ERROR: Expected: %v, got: %v", jt.firstResult, jsonString)
		} else {
			cars := fromJson(jsonString)
			cars = append(cars, jt.c1)
			cars = append(cars, jt.c2)
			cars = append(cars, jt.c3)
			jsonStringAfterAppend := toJson(cars)
			if jt.secondResult != jsonStringAfterAppend {
				t.Errorf("ERROR: Expected: %v, got: %v", jt.secondResult, jsonStringAfterAppend)
			}
		}
	}
}