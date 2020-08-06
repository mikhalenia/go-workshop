package queue

import (
	"errors"
	"job"
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
	Push(job.IJob)
	Pop() (job.IJob, error)
}

type Queue struct {
	sync.Mutex
	Content []job.IJob
}

func (q *Queue) Push(j job.IJob) {
	q.Lock()
	q.Content = append(q.Content, j)
	q.Unlock()
}

func (q *Queue) Pop() (job.IJob, error) {
	if len(q.Content) < 1 {
		return nil, errors.New("empty queue")
	}
	q.Lock()
	j := q.Content[0]
	if len(q.Content) > 0 { // If we get not last job
		q.Content = q.Content[1:]
	} else {
		q.Content = []job.IJob{} // Clear list
	}
	q.Unlock()
	return j, nil
}
