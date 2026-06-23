package main

import (
	"fmt"
	"distribute-job-queue/src/queue"
)

func main() {
	q := queue.NewQueue(10)
	msg := queue.NewMessage("Hello there")
	q.Enqueue(msg)
	fmt.Print(q)
}