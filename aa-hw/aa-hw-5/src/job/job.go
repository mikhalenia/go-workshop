package job

import (
	"fmt"
)

type IJob interface {
	Working() string
}

type DecJob struct {
	J int
}

type StrJob struct {
	J string
}

func (j DecJob) Working() string {
	return fmt.Sprintf("decimal %d", j.J)
}

func (j StrJob) Working() string {
	return fmt.Sprintf(j.J)
}
