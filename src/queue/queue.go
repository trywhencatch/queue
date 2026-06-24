package queue


// Ring Buffer

// essentially we need a array that wraps itself and we need pointers on head and tail to essentially add and remove elments not by actually
// remving but by moving the reference pointer and assuming its gone as we will have fixed ring buffer capacity when the tail is at its last
// index it will wrap back then our head will be somewhere where the messages start from we will have this space for entries that we have
// assumed to be gone


import (
	"fmt"
)

type Message struct {
	msg string
}

func NewMessage(msg string) Message {
	return Message{
		msg: msg,
	}
}

type Queue struct {
	data     []Message
	head     int
	tail     int
	size     int
	capacity int
}

func NewQueue(capacity int) Queue {
	return Queue{
		data:     make([]Message, capacity),
		capacity: capacity,
	}
}

func (q *Queue) Enqueue(msg Message) error {
	if q.size == q.capacity {
		return fmt.Errorf("queue capacity reached")
	}
	q.data[q.tail] = msg
	q.size++
	q.tail = (q.tail + 1) % q.capacity
	fmt.Println("Message added to queue")
	return nil
}

func (q *Queue) Dequeue() (Message, error) {
	if q.size == 0 {
		return Message{}, fmt.Errorf("queue empty")
	}
	returnVal := q.data[q.head]
	q.size--
	q.head = (q.head + 1) % q.capacity
	return returnVal, nil
}

func (q *Queue) String() string {
	return fmt.Sprintf("messages: %v, size:%d, tail: %d}", q.data, q.size, q.tail)
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
