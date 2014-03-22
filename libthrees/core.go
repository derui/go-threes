package libthrees

import (
	"math"
	"strconv"
)

type three struct {
	value int
	base int
}

func empty() three {
	return three{0, 0}
}

func (t three) IsEmpty() bool {
	return t.value == 0 && t.base == 0
}

func GetThree(num int) three {
	if num <= 0 {
		return empty()
	}

	switch num {
	case 1:
		fallthrough;
	case 2:
		return three{num, 0}
	default:
		break;
	}

	return three{3 * int(math.Floor(math.Pow(2., float64(num - 3)))), num - 3}
}

func (t three) String() string {
	return strconv.Itoa(t.value)
}

func (t three) Value() int {
	return t.value
}

func (t three) Score() int {

	if t.IsEmpty() {
		return 0
	}

	switch t.value {
	case 1:
		fallthrough
	case 2:
		return 0
	default:
		break
	}

	pow := math.Floor(math.Pow(3., float64(t.base + 1)))
	return int(pow)
}

func (t three) CanMerge(other three) bool {
	switch t.value {
	case 1:
		return other.value == 2
	case 2:
		return other.value == 1
	default:
		return t.value == other.value
	}
}







