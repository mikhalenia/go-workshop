package job

import (
	"errors"
	"sync"
)

type NewQueue struct {
	sync.Mutex
	Content []string
}

func (q *NewQueue)Push(s interface{}) {
	q.Lock()
	q.Content = append(q.Content, s.(string))
	q.Unlock()
}

func (q *NewQueue)Pop() (interface{}, error) {
	if len(q.Content) < 1 {
		return "", errors.New("empty NewQueue")
	}
	s := q.Content[0]
	q.Lock()
	q.Content = q.Content[1:]
	q.Unlock()
	return s, nil
}
