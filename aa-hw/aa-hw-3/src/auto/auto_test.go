package auto

import (
	"fmt"
	"os"
	"testing"
	//"time"
)

var autos Catalog

var auto Auto

func TestAutoCRUD(t *testing.T) {
	autos = CreateCatalog()

	// Create
	for i := 1; i <= 10; i++ {
		auto = Auto{autos.createId(), fmt.Sprintf("mark %d", i), fmt.Sprintf("model %d", i), fmt.Sprintf("url %d", i), i, "free"}
		autos.AddAuto(auto)
	}

	// Read
	a := autos.GetAuto(5)

	// Update
	a.Status = "busy"
	autos.UpdateAuto(a)
	if autos.Table[5].Status != "busy" {
		t.Errorf("Error Test Update Auto. Want = busy and Got = '%s'.", autos.Table[5].Status)
	}

	//Delete
	autos.DeleteAuto(a.Id)

	auto = Auto{autos.createId(), fmt.Sprintf("mark %d", 0), fmt.Sprintf("model %d", 0), fmt.Sprintf("url %d", 0), 0, "busy"}
	autos.AddAuto(auto)

	a = autos.GetAuto(5)

	if a.Id != 5 {
		t.Errorf("Error Test Number Auto. Want = 5 and Got = '%d'.", a.Id)
	}

	if len(autos.Table) != 10 {
		t.Errorf("Error Test Len Autos. Want = 10 and Got = '%d'.", len(autos.Table))
	}
	auto = Auto{autos.createId(), fmt.Sprintf("mark %d", 11), fmt.Sprintf("model %d", 11), fmt.Sprintf("url %d", 11), 11, "free"}
	autos.AddAuto(auto)
	a = autos.GetAuto(len(autos.Table))

	if a.Id != 11 {
		t.Errorf("Error Test Number Auto. Want = 11 and Got = '%d'.", a.Id)
	}

	if len(autos.Table) != 11 {
		t.Errorf("Error Test Len Autos. Want = 11 and Got = '%d'.", len(autos.Table))
	}

}

func TestAutoLoad(t *testing.T) {
	var err error
	var f *os.File

	filePath := "../test_auto.json"
	autos = CreateCatalog()

	deleteFile := func(f *os.File, filePath string) {
		//time.Sleep(60 * time.Second) // wait for can look file from FS
		f.Close()
		if err := os.Remove(filePath); err != nil {
			t.Error("Error Test AutoLoad. Can't delete test file.")
		}
	}

	// Open or create test file
	if f, err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0660); err != nil {
		if f, err = os.Create(filePath); err != nil {
			t.Error("Error Test AutoLoad. Can't open or create test file")
			return
		} else {
			t.Log("Test File Created")
		}
	}

	defer deleteFile(f, filePath)

	AutoArray := []string{
		`{"id":1,"mark":"fiat","model":"punto2","url":"","count":60,"status":"free"}`,
		`{"id":2,"mark":"fiat","model":"scudo","url":"","count":120,"status":"free"}`,
		`{"id":3,"mark":"alfaromeo","model":"420","url":"","count":80,"status":"busy"}`,
		`{"id":4,"mark":"luncia","model":"cupa","url":"","count":40,"status":"free"}`,
	}

	count := 0 // Count Auto
	// Insert test Auto to file
	for _, str := range AutoArray {
		if _, err = f.WriteString(str + "\n"); err != nil {
			t.Error("Error Test AutoLoad. Can't insert test Autos in test file")
			return
		}
		count++
	}

	f.Close()

	// Load auto from file to Catalog
	if err = autos.LoadAuto(filePath); err != nil {
		t.Error("Error Test AutoLoad. Can't load Autos from test file")
	}

	// View auto from Catalog
	autoList := autos.GetAll()
	for i, a := range autoList {
		t.Logf("Result Test AutoLoad. Auto #%d: %v.", i, a)
		count--
	}
	if count != 0 {
		t.Error("Error Test AutoLoad. Wrong count loaded auto")
	}

}
