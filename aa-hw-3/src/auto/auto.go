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

type Stack struct {
	sync.Mutex
	Table map[int]Auto
}

func CreateStack() Stack {
	return Stack{
		Table: make(map[int]Auto),
	}
}

func (s *Stack) SetAuto(value ...interface{}) error { // got some value for create Auto
	auto := Auto{}
	auto.id = s.createId()

	for i, val := range value {
		switch i {
		case 0:
			auto.mark = val.(string)
		case 1:
			auto.model = val.(string)
		case 2:
			auto.url = val.(string)
		case 3:
			auto.count = val.(int)
		}
	}
	logs.Logs(fmt.Sprintf("Created auto: %v", auto))
	s.Lock()
	s.Table[auto.id] = auto
	s.Unlock()
	return nil
	return nil
}

func (s *Stack) GetAuto(id int) Auto {
	return s.Table[id]
}

func (s *Stack) DeleteAuto(id int) {
	s.Lock()
	delete(s.Table, id)
	s.Unlock()
}

func (s *Stack) createId() (id int) {
	i := 1
	for {
		if _, ok := s.Table[i]; !ok {
			id = i
			logs.Logs(fmt.Sprintf("Created id: %d", id))
			return
		}
		i++
	}
}
