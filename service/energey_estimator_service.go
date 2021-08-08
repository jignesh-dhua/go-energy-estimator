package service

import "github.com/jignesh-dhua/go-energy-estimator/message"

func EstimateEnergyConsumption(messages []message.Message) float64 {

	return 0
}

func CalculateDuration(current message.Message, next message.Message) int64 {

	if next.IsEmpty() {
		return next.Timestamp.Unix() - current.Timestamp.Unix()
	} else {
		return 0
	}

}
