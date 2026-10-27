package comm

import (
	"math"
)

var DropRate uint8 = 1

func GetXpRequired(lvl uint8, xpScalingFactor float64) float64 {
	return 50 * math.Pow(xpScalingFactor, float64(lvl-1))
}
