package broker

// Ring Buffer

// essentially we need a array that wraps itself and we need pointers on head and tail to essentially add and remove elments not by actually
// remving but by moving the reference pointer and assuming its gone as we will have fixed ring buffer capacity when the tail is at its last
// index it will wrap back then our head will be somewhere where the messages start from we will have this space for entries that we have
// assumed to be gone

import (
	"fmt"
	"sync"
)

type Queue struct {
	data     []Message
	head     int
	tail     int
	size     int
	capacity int
	mu sync.Mutex
	cond sync.Cond
}

func NewQueue(capacity int) *Queue {
	q := &Queue{
		data:     make([]Message, capacity),
		capacity: capacity,
	}

	q.cond = *sync.NewCond(&q.mu)
	return q
}

func (q *Queue) Enqueue(msg Message) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.data[q.tail] = msg
	q.size++
	q.tail = (q.tail + 1) % q.capacity

	q.cond.Signal()
}

func (q *Queue) Dequeue() (Message, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	for q.size == 0 {
		q.cond.Wait()
	}
	returnVal := q.data[q.head]
	q.size--
	q.head = (q.head + 1) % q.capacity
	return returnVal, nil
}

func (q *Queue) String() string {
	return fmt.Sprintf("messages: %v, size:%d, tail: %d", q.data, q.size, q.tail)
}

func (q *Queue) Peek() (Message, error) {
	return q.data[q.head], nil
}

func (q *Queue) Len() int {
	return q.size
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) IsFull() bool {
	return q.size == q.capacity
}
