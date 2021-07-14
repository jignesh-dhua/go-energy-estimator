package service

import "math"

func calculate(currentvalue float64, previousvalue float64) float64 {

	if currentvalue == 0 {
		return 0
	} else {

		consumption := math.Abs(previousvalue) + currentvalue

		if consumption < 0 {
			return 0
		}

		return 5 * (math.Abs(previousvalue) + currentvalue)
	}
}
