package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jignesh-dhua/go-energy-estimator/message"
	"github.com/jignesh-dhua/go-energy-estimator/service"
)

func main() {

	var messages = ReadMessages()

	fmt.Println("*** Messages ***")
	fmt.Println(messages)
}

func ReadMessages() []message.Message {
	scanner := bufio.NewScanner(os.Stdin)
	//scanner.Split(bufio.ScanLines)

	var messages []message.Message

	for scanner.Scan() {
		_, lines, err := bufio.ScanLines([]byte(scanner.Text()), true)
		if err == nil {
			fmt.Println("line :: ", string(lines))

			messages = append(messages, service.ParseFrom(string(lines)))
		}
	}

	return messages
}
