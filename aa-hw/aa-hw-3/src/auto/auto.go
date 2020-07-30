package auto

import (
	"bufio"
	"encoding/json"
	"fmt"
	"logs"
	"os"
	"strings"
	"sync"
)

type AutoDB interface {
	GetAll() []Auto
	AddAuto(Auto) error
	UpdateAuto(Auto) error
	GetAuto(int) Auto
	DeleteAuto(int)
}

type Auto struct {
	Id     int    `json:"id"`
	Brend  string `json:"brend"`
	Model  string `json:"model"`
	URL    string `json:"url"`
	Count  int    `json:"count"`
	Status string `json:"status"`
}

type Catalog struct {
	sync.Mutex
	Table map[int]Auto
}

func CreateCatalog() Catalog {
	return Catalog{
		Table: make(map[int]Auto),
	}
}

func (c *Catalog) AddAuto(auto Auto) error { // got some value for create Auto
	logs.Logs(fmt.Sprintf("Created auto: %v", auto))
	c.Lock()
	c.Table[auto.Id] = auto
	c.Unlock()
	return nil
}

func (c *Catalog) UpdateAuto(auto Auto) error {
	return c.AddAuto(auto)
}

func (c *Catalog) GetAuto(id int) Auto {
	return c.Table[id]
}

func (c *Catalog) DeleteAuto(id int) {
	c.Lock()
	delete(c.Table, id)
	c.Unlock()
}

func (c *Catalog) GetAll() []Auto {
	result := []Auto{}
	c.Lock()
	for _, a := range c.Table {
		result = append(result, a)
	}
	c.Unlock()
	return result
}

func (c *Catalog) LoadAuto(file string) error {

	f, err := os.Open(file)
	if err != nil {
		return err
	}

	defer f.Close()
	var a Auto

	fileData := bufio.NewScanner(f)
	for fileData.Scan() {
		dec := json.NewDecoder(strings.NewReader(fileData.Text()))
		if err = dec.Decode(&a); err != nil {
			return err
		}
		c.Table[a.Id] = a
	}
	if err := fileData.Err(); err != nil {
		return err
	}
	return nil
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
