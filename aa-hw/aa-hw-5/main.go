package main

import (
	"fmt"
	"time"

	"job"
	"queue"
)

func Worker(id int, q *queue.Queue, quit chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			if x, err := q.Pop(); err == nil {
				fmt.Printf("Worker #%d do %s work\n", id, x.Working())
			} else {
				fmt.Printf("Worker #%d, waiting work: %s\n", id, err.Error())
				time.Sleep(time.Second)
			}
		}
	}
}

func Broker(q *queue.Queue, quit chan bool) {
	x := 0
	var j job.IJob
	for {
		select {
		case <-quit:
			return
		default:
			if x%2 > 0 {
				j = job.DecJob{x}
				q.Push(j)
			} else {
				j = job.StrJob{fmt.Sprintf("string %d", x)}
				q.Push(j)
			}
			fmt.Printf("Pushed Job: %v\n", j)
			x++
			if x > 100 {
				time.Sleep(time.Second)
				x = 0
			}
		}
	}
}

func main() {

	messages := new(queue.Queue)

	quitWorker := make(chan bool)
	quitBroker := make(chan bool)

	for i := 0; i < 100; i++ {
		go Worker(i, messages, quitWorker)
	}

	go Broker(messages, quitBroker)

	time.Sleep(time.Second * 10)

	// kill Broker
	close(quitBroker)

	for len(messages.Content) > 0 {
	}

	// kill workers
	close(quitWorker)

	fmt.Println("Done")
}
