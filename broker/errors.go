package broker

import "errors"

var (
	ErrQueueEmpty = errors.New("queue is empty")
)