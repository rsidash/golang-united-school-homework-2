package square

import (
	"math"
)

const SidesTriangle int = 3
const SidesSquare int = 4
const SidesCircle int = 0

type sidesNumType int

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
