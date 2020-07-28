package job

import "fmt"

type Job interface {
	PrintInfo() string
}

type JobIndex struct {
	Index int
}

func (j JobIndex) PrintInfo() string {
	return fmt.Sprintf(" job with index : %d", j.Index)
}