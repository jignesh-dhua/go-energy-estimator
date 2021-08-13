package service

import (
	"math"

	"github.com/jignesh-dhua/go-energy-estimator/message"
)

func EstimateEnergyConsumption(messages []message.Message) float64 {

	var totalConsumption float64 = 0

	for i := 1; i < (len(messages) - 1); i++ {

		var previous = messages[i-1]
		var current = messages[i]
		var next = messages[i+1]

		durationInSeconds := CalculateDuration(current, next)
		durationInHr := durationInSeconds / 3600

		consumption := Calculate(current.Value, previous.Value)

		//totalConsumption += Math.abs(durationInHr * consumption)

		totalConsumption += math.Abs(float64(durationInHr) * consumption)

	}

	return totalConsumption
}

func CalculateDuration(current message.Message, next message.Message) int64 {

	if !next.IsEmpty() {
		return next.Timestamp.Unix() - current.Timestamp.Unix()
	} else {
		return 0
	}

}
