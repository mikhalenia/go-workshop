package queue

import "fmt"

type Job interface {
	RunJob() string
}

type JobMessageEven struct {
	MessageEven string
}

type JobMessageOdd struct {
	MessageOdd string
}

func (j JobMessageEven) RunJob() string {
	return fmt.Sprintf("%d", j.MessageEven)
}

func (j JobMessageOdd) RunJob() string {
	return fmt.Sprintf("%d", j.MessageOdd)
}
