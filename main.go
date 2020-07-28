package main

import (
	"./aa-hw/queue"
	"fmt"
	"time"
)



func Worker(id int, q *queue.Queue, quit chan bool) {
	for {
		select {
		case <- quit:
			return
		default:
			if x, err := q.Pop(); err == nil {
				fmt.Printf("Worked id #%d, receiving %v\n", id, x)
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
			q.Push(x)
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
		if i%2>0 {
			go Worker(i, messages, quitWorker)
		}
	}

	go Broker(messages, quitBroker)

	time.Sleep(time.Second * 10)

	// kill Broker
	close(quitBroker)

	for len(messages.Content) > 0 {}

	// kill workers
	close(quitWorker)

	fmt.Println("Done")
}
