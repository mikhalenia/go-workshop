package main

import (
	"./queue"
	"./job"
	"fmt"
	"time"
)

func Worker(id int, i queue.IQueue, quit chan bool) {
	for {
		select {
		case <- quit:
			return
		default:
			if x, err := i.Pop(); err == nil {
				fmt.Printf("Worked id #%d, receiving %v\n", id, x)
			} else {
				fmt.Printf("Worker id #%d, Error: %s, need to wait\n", id, err.Error())
				time.Sleep(time.Second)
			}
		}
	}
}

func Broker(i queue.IQueue, quit chan bool) {
	x := 0
	for {
		select {
		case <- quit:
			return
		default:
			//switch i.(type) {
			switch i {
			case iMessages:
				i.Push(x)
			case iNewJob:
				i.Push(fmt.Sprintf("Broker %d",x))
			}
			x++
			if x > 100 {
				time.Sleep(time.Second)
				x = 0
			}
		}
	}
}

var (
	iMessages queue.IQueue
	iNewJob   queue.IQueue
)

func main() {

	messages := new(queue.Queue)
	newJobs := new(job.NewQueue)

	iMessages = queue.IQueue(messages)
	iNewJob = queue.IQueue(newJobs)

	quitWorker := make(chan bool)
	quitBroker := make(chan bool)

	for i := 0; i < 100; i++ {
		if i%2>0 {
			go Worker(i, iMessages, quitWorker)
		} else {
			go Worker(i, iNewJob, quitWorker)
		}
	}

	go Broker(iMessages, quitBroker)
	go Broker(iNewJob, quitBroker)

	time.Sleep(time.Second * 10)

	// kill Broker
	close(quitBroker)

	for len(messages.Content) > 0 {}

	// kill workers
	close(quitWorker)

	fmt.Println("Done")
}
