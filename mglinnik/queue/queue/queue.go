package queue
import (
	"errors"
	"sync"
)

type Queue struct {
	Content []Job
	mux sync.Mutex
}
func (q *Queue) Push(j Job) {
	q.mux.Lock()
	q.Content = append(q.Content, j)
	defer q.mux.Unlock()
}
func (q *Queue) Pop() (Job, error) {
	if len(q.Content) < 1 {
		return nil, errors.New("empty queue")
	}
	q.mux.Lock()
	i := q.Content[0]
	q.Content = q.Content[1:]
	defer q.mux.Unlock()
	return i, nil
}