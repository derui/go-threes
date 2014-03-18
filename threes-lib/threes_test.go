package threeslib

import (
	"testing"
	"fmt"
)

var threeMap map[int]int = map[int]int{
	0 : 1,
	1 : 1,
	2 : 2,
	3 : 3,
	4 : 6,
	5 : 12,
	6 : 24,
	7 : 48,
	8 : 96,
	9 : 192,
	10 : 384,
	11 : 768,
	12 : 1536,
}

func TestCanGetThreesWithNum(t *testing.T) {

	for k, v := range(threeMap) {
		if GetThree(k).value != v {
			t.Errorf("Getting threes value with %d should be %d, but getting %d",
				k, v, GetThree(k))
		}
	}
}

func TestCanThreesToString(t *testing.T) {

	for k, v := range(threeMap) {
		tv := GetThree(k)
		if fmt.Sprint(tv) != fmt.Sprintf("%d", v) {
			t.Errorf("Getting threes value with %d should be %s, but getting %s",
				k, fmt.Sprintf("%d", v), fmt.Sprint(tv))
		}
	}
}
