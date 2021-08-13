package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jignesh-dhua/go-energy-estimator/message"
	"github.com/jignesh-dhua/go-energy-estimator/service"
)

func main() {

	// var messages = []string{"1544206562 TurnOff",
	// 	"1544206563 Delta +0.5",
	// 	"1544210163 TurnOff"}

	var messages = ReadMessages()
	fmt.Println("*** Messages ***")
	fmt.Println(messages)

	estimation := service.EstimateEnergyConsumption(message.Map(messages, service.ParseFrom))
	fmt.Println("Estimation :: ", estimation)
}

func ReadMessages() []string {
	scanner := bufio.NewScanner(os.Stdin)
	//scanner.Split(bufio.ScanLines)

	var messages []string

	for scanner.Scan() {
		_, lines, err := bufio.ScanLines([]byte(scanner.Text()), true)
		if err == nil {
			fmt.Println("line :: ", string(lines))

			messages = append(messages, string(lines))
		}
	}

	return messages
}
