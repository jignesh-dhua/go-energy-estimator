package main

import (
	"fmt"
	"time"

	"github.com/jignesh-dhua/go-energy-estimator/message"
	"github.com/jignesh-dhua/go-energy-estimator/service"
)

func main() {
	fmt.Println("Hello World")
	messages := []message.Message{

		{Timestamp: time.Now(), MessageType: "Test", Value: 0.0},
	}

	for i := 0; i < len(messages); i++ {
		fmt.Println(messages[i].Timestamp)
	}

	message := service.ParseFrom("1544206563 Delta -0.5")

	fmt.Println(message)
}
