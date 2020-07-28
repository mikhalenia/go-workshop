package queue

import (
	"errors"
	"studying_go_high_technologies_park/homework_5/job"
	"sync"
)

type Queue struct {
	mutex sync.Mutex
	Content []job.Job
}

func (q *Queue)Push(i job.Job) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.Content = append(q.Content, i)
}

func (q *Queue) Pop() (job.Job, error) {
	if len(q.Content) < 1 {
		return nil, errors.New("empty queue")
	}
	q.mutex.Lock()
	i := q.Content[0]
	q.Content = q.Content[1:]
	defer q.mutex.Unlock()
	return i, nil
}