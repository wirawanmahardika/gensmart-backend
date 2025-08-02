package pkg

import "math"

func RoundToTwoDecimal(x float64) float64 {
	return math.Round(x*100) / 100
}
