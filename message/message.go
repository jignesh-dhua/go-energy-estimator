package message

import (
	"reflect"
	"time"
)

type Message struct {
	Timestamp   time.Time
	MessageType string
	Value       float64
}

func (m Message) IsEmpty() bool {
	return reflect.DeepEqual(m, Message{})
}

type messages []Message

func Map(msgs []string, f func(string) Message) messages {

	messages := make([]Message, len(msgs))

	for i, m := range msgs {
		messages[i] = f(m)
	}

	return messages
}
