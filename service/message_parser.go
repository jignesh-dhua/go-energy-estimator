package service

import (
	"strconv"
	"strings"
	"time"

	"github.com/jignesh-dhua/go-energy-estimator/message"
)

func ParseFrom(msg string) message.Message {

	var message message.Message

	fields := strings.Fields(msg)

	// fmt.Println(fields[0])
	// fmt.Println(fields[1])

	timeInMilli, _ := strconv.ParseInt(fields[0], 10, 64)

	message.Timestamp = time.Unix(timeInMilli, 0)

	switch message.MessageType = fields[1]; message.MessageType {

	case "Delta":
		{
			//fmt.Println("Message type is delta")

			if delta, err := strconv.ParseFloat(fields[2], 64); err == nil {
				message.Value = delta
			}

		}
	case "TurnOff":
		{
			//fmt.Println("Message type is turnoff")
		}
	default:
		{
			// message type not supported
		}
	}

	//fmt.Println(message)

	return message
}
