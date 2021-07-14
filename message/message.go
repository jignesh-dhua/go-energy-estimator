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
