package queue

import (
	"errors"
	"sync"

	"go-workshop/frogstair/hw5/jobs"
)

// Queue represent a queue
type Queue struct {
	Content []jobs.Job
	mux     sync.Mutex
}

// Push pushes an object onto the queue
func (q *Queue) Push(j jobs.Job) {
	q.Content = append(q.Content, j)
}

// Pop removes an object from the queue
func (q *Queue) Pop() (jobs.Job, error) {
	if len(q.Content) < 1 {
		return nil, errors.New("empty queue")
	}
	q.mux.Lock()
	i := q.Content[0]
	q.Content = q.Content[1:]
	defer q.mux.Unlock()
	return i, nil
}
