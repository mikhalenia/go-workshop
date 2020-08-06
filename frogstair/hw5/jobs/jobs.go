package jobs

import "fmt"

// Job is a job that a worker would do
type Job interface {
	Run() string
}

// PrintJob is a print job
type PrintJob struct {
	Text string
}

// NumberJob adds two numbers and prints it
type NumberJob struct {
	A, B int
}

// Run runs a job
func (p PrintJob) Run() string {
	return fmt.Sprintf(p.Text)
}

// Run runs a job
func (p NumberJob) Run() string {
	return fmt.Sprintf("%d", p.A+p.B)
}
