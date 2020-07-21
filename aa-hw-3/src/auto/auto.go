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

type List struct {
	sync.Mutex
	Table map[int]Auto
}

func CreateList() List {
	return List{
		Table: make(map[int]Auto),
	}
}

func (l *List) SetAuto(auto Auto) error { // got some value for create Auto
	//	auto := Auto{}
	//	auto.id = l.createId()

	//	for i, val := range value {
	//		switch i {
	//		case 0:
	//			auto.mark = val.(string)
	//		case 1:
	//			auto.model = val.(string)
	//		case 2:
	//			auto.url = val.(string)
	//		case 3:
	//			auto.count = val.(int)
	//		}
	//	}
	logl.Logs(fmt.Sprintf("Created auto: %v", auto))
	l.Lock()
	l.Table[auto.id] = auto
	l.Unlock()
	return nil
	return nil
}

func (l *List) GetAuto(id int) Auto {
	return l.Table[id]
}

func (l *List) DeleteAuto(id int) {
	l.Lock()
	delete(l.Table, id)
	l.Unlock()
}

func (l *List) createId() (id int) {
	i := 1
	for {
		if _, ok := l.Table[i]; !ok {
			id = i
			logl.Logs(fmt.Sprintf("Created id: %d", id))
			return
		}
		i++
	}
}
