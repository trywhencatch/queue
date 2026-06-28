package broker

import (
	"fmt"
)

type Message struct {
	Msg string
}

func NewMessage(msg string) Message {
	return Message{
		Msg: msg,
	}
}

func (m *Message) String() string {
	return fmt.Sprintf("message: %v", m.Msg)
}