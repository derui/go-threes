package threeslib

import (
	"math"
	"fmt"
)

type three struct {
	value int
}

func GetThree(num int) three {
	if num <= 0 {
		return three{1}
	}

	switch num {
	case 1:
		return three{num}
	case 2:
		return three{num}
	default:
		break;
	}

	return three{3 * int(math.Floor(math.Pow(2., float64(num - 3))))}
}

func (t three) String() string {
	return fmt.Sprintf("%d", t.value)
}
