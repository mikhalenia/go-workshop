package main

import (
	"fmt"
	"math/rand"
	"studying_go_high_technologies_park/homework_5/job"
	"studying_go_high_technologies_park/homework_5/queue"
	"time"
)

func Worker(id int, q *queue.Queue, quit chan bool) {
	for {
		select {
		case <- quit:
			return
		default:
			if x, err := q.Pop(); err == nil {
				fmt.Printf("Worked id #%d, receiving : %s\n", id, x.PrintInfo())
			} else {
				fmt.Printf("Worker id #%d, Error: %s, need to wait\n", id, err.Error())
				time.Sleep(time.Second)
			}
		}
	}

}

func Broker(q *queue.Queue, quit chan bool) {
	x := 0
	for {
		select {
		case <- quit:
			return
		default:
			min := 1000000
			max := 9999999
			q.Push(job.JobIndex{Index: rand.Intn(max - min) + min})
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

	close(quitBroker)

	for len(messages.Content) > 0 {}

	close(quitWorker)

	fmt.Println("Done")
}