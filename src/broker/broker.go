package broker

import (
	"distribute-job-queue/src/queue"
	"fmt"
)

type Broker struct {
	queues map[string]queue.Queue
}

func NewBroker() *Broker {
	return &Broker{
		queues: make(map[string]queue.Queue),
	}
}

func (b *Broker) CreateQueue(queueName string) {
	q := queue.NewQueue(10)
	b.queues[queueName] = q
}

func (b *Broker) DeleteQueue(queueName string) {
	delete(b.queues, queueName)
}

func (b *Broker) Publish(queueName string, msg queue.Message) error {
	q, exists := b.queues[queueName]
	if !exists {
		return fmt.Errorf("Queue with name %s does not exist", queueName)
	}
	return q.Enqueue(msg)
}

func (b *Broker) Consume(queueName string) (queue.Message, error) {

	q, exists := b.queues[queueName]

	if !exists {
		return queue.Message{}, fmt.Errorf("Queue with name %s does not exist", queueName)
	}

	msg, err := q.Dequeue()

	if err != nil {
		return queue.Message{}, fmt.Errorf("Queue Empty")
	}

	return msg, nil
}

func (b *Broker) ListQueues() []string {
	queues := make([]string, len(b.queues))

	for name := range b.queues {
		queues = append(queues, name)
	}

	return queues
}