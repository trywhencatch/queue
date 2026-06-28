package broker

import (
	"bufio"
	"fmt"
	"net"
)

type Exchange struct {
	queues map[string]*Queue
}

func NewExchange() *Exchange {
	return &Exchange{
		queues: make(map[string]*Queue),
	}
}

func (e *Exchange) CreateQueue(queueName string) {
	q := NewQueue(10)
	e.queues[queueName] = q
}

func (e *Exchange) DeleteQueue(queueName string) {
	delete(e.queues, queueName)
}

func (e *Exchange) Publish(queueName string, msg Message) error {
	q, exists := e.queues[queueName]
	if !exists {
		return fmt.Errorf("Queue with name %s does not exist", queueName)
	}
	q.Enqueue(msg)
	return nil
}

func (e *Exchange) GetQueue(queueName string) (*Queue, error) {
    q, exists := e.queues[queueName]
    if !exists {
        return nil, fmt.Errorf("Queue with name %s does not exist", queueName)
    }
    return q, nil
}

func (e *Exchange) Consume(queueName string) (Message, error) {
	q, exists := e.queues[queueName]
	if !exists {
		return Message{}, fmt.Errorf("Queue with name %s does not exist", queueName)
	}
	msg, err := q.Dequeue()
	if err != nil {
		return Message{}, fmt.Errorf("Queue Empty")
	}
	return msg, nil
}

func (e *Exchange) ListQueues() []string {
	queues := make([]string, len(e.queues))

	for name := range e.queues {
		queues = append(queues, name)
	}

	return queues
}

func (e *Exchange) HandleProducer(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	success := false
	for {
		data, err := reader.ReadBytes('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Client disconnected")
				fmt.Println(e.queues)
				return
			}
			fmt.Print(
				"Failed to read data", err,
			)
			return
		}

		msg := NewMessage(string(data))
		err = e.Publish("test_queue", msg)

		if err != nil {
			fmt.Println("Failed to push to queue")
			conn.Write([]byte("unack"))
			break
		}

		success = true
	}
	if success {
		_, err := conn.Write([]byte("ack\n"))
		if err != nil {
			fmt.Println("Write failed or timed out:", err)
			conn.Close()
			return
		}
	}
}

func (e *Exchange) HandleConsumer(conn net.Conn) {
	msg, err := e.Consume("test_queue")
	if err != nil {
		fmt.Printf("Error occured trying to consume queue %v %v", "test_queue \n", err)
		return
	}
	byteArr := []byte(msg.Msg)
	_,err = conn.Write(byteArr)
	if err != nil {
		fmt.Println("Error Writing to the consumer")
		conn.Close()
	}
}
