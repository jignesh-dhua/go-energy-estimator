package main

import (
	"fmt"
	"time"

	"github.com/jignesh-dhua/go-energy-estimator/message"
)

func main() {
	fmt.Println("Hello World")
	messages := []message.Message{

		{Timestamp: time.Now(), MessageType: "Test", Value: 0.0},
	}

	for i := 0; i < len(messages); i++ {
		fmt.Println(messages[i].Timestamp)
	}
}
