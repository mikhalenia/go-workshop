package main

import (
	"fmt"
	"math/rand"
	"time"

	"go-workshop/frogstair/hw5/jobs"
	"go-workshop/frogstair/hw5/queue"
)

func worker(id int, q *queue.Queue, quit chan bool) {
	for {
		select {
		case <-quit:
			return
		default:
			if x, err := q.Pop(); err == nil {
				fmt.Printf("Worker #%d: Ran and got result %q\n", id, x.Run())
			} else {
				fmt.Printf("Worker %d: No job yet\n", id)
			}
		}
		time.Sleep(time.Second * time.Duration(rand.Intn(3)+1))
	}
}

func broker(q *queue.Queue, quit chan bool) {
	x := 0
	for {
		select {
		case <-quit:
			return
		default:
			if randBool() {
				q.Push(jobs.PrintJob{Text: randStr(10)})
			} else {
				q.Push(jobs.NumberJob{A: rand.Int() % 50, B: rand.Int() % 50})
			}
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
		go worker(i, messages, quitWorker)
	}

	go broker(messages, quitBroker)

	time.Sleep(time.Second * 10)

	// kill Broker
	close(quitBroker)

	for len(messages.Content) > 0 {
	}

	// kill workers
	close(quitWorker)

	fmt.Println("Done")
}

func randBool() bool {
	return rand.Float32() < 0.5
}

func randStr(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
