package queue

import (
	"errors"
	"sync"
)

// queue
// FIFO
// []  push 1
// [1] push 2
// [1,2] pop -> 1
// [1] pop -> 2
// [] pop -> nil


type IQueue interface {
	Push(interface{})
	Pop() (interface{},error)
}

type Queue struct {
	sync.Mutex
	Content []int
}

func (q *Queue)Push(i interface{}) {
	q.Lock()
	q.Content = append(q.Content, i.(int))
	q.Unlock()
}

func (q *Queue)Pop() (interface{}, error) {
	if len(q.Content) < 1 {
		return 0, errors.New("empty queue")
	}
	i := q.Content[0]
	q.Lock()
	q.Content = q.Content[1:]
	q.Unlock()
	return i, nil
}