package auto

import (
	"fmt"
	"logs"
	"sync"
)

type Auto struct {
	id     int
	mark   string
	model  string
	url    string
	count  int
	status bool
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

func (c *Catalog) SetAuto(auto Auto) error { // got some value for create Auto
	//	auto := Auto{}
	//	auto.id = c.createId()

	//	for i, val := range value {
	//		switch i {
	//		case 0:
	//			auto.mark = vac.(string)
	//		case 1:
	//			auto.model = vac.(string)
	//		case 2:
	//			auto.url = vac.(string)
	//		case 3:
	//			auto.count = vac.(int)
	//		}
	//	}
	logs.Logs(fmt.Sprintf("Created auto: %v", auto))
	c.Lock()
	c.Table[auto.id] = auto
	c.Unlock()
	return nil
	return nil
}

func (c *Catalog) GetAuto(id int) Auto {
	return c.Table[id]
}

func (c *Catalog) DeleteAuto(id int) {
	c.Lock()
	delete(c.Table, id)
	c.Unlock()
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
