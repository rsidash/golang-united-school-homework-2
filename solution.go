package square

import (
	"math"
)

type sidesNumType int

const SidesTriangle sidesNumType = 3
const SidesSquare sidesNumType = 4
const SidesCircle sidesNumType = 0

func CalcSquare(sideLen float64, sidesNum sidesNumType) float64 {

	switch sidesNum {
	case 3:
		return sideLen * sideLen * math.Sqrt(3) / 4
	case 4:
		return sideLen * sideLen * sideLen
	case 0:
		return math.Pi * sideLen * sideLen
	default:
		return 0
	}

}
