package libthrees

import (
	"testing"
	"fmt"
)


var emptyMap = map[int]bool{
	0 : true,
	1 : false,
	2 : false,
	3 : false,
	4 : false,
	5 : false,
	6 : false,
	7 : false,
	8 : false,
	9 : false,
	10 : false,
	11 : false,
	12 : false,
}

func TestCanDetectEmptyThree(t *testing.T) {

	for k,v := range(emptyMap) {
		emp := GetThree(k)

		if emp.IsEmpty() != v {
			t.Errorf("The empty of three is returned GetThree with lesser than 0")
		}
	}
}

var threeMap = map[int]int{
	0 : 0,
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
		if GetThree(k).Value() != v {
			t.Errorf("Getting threes value with %d should be %d, but getting %d",
				k, v, GetThree(k))
		}
	}
}

func TestCanThreesToString(t *testing.T) {

	for k, v := range(threeMap) {
		tv := GetThree(k)
		if fmt.Sprint(tv) != fmt.Sprintf("%d", v) {
			t.Errorf("Getting threes as string with %d should be %s, but getting %s",
				k, fmt.Sprintf("%d", v), fmt.Sprint(tv))
		}
	}
}

var scoreMap = map[int]int{
	0 : 0,
	1 : 0,
	2 : 0,
	3 : 3,
	4 : 9,
	5 : 27,
	6 : 81,
	7 : 243,
	8 : 729,
	9 : 2187,
	10 : 6561,
	11 : 19683,
	12 : 59049,
}


func TestCanGetScoreOfAThree(t *testing.T) {

	for k,v := range(scoreMap) {
		tv := GetThree(k)

		if tv.Score() != v {
			t.Errorf("Getting score of threes with %d should be %d, but getting %d",
				k, v, tv.Score())
		}
	}
}

func TestCanMergeBetweenThrees(t *testing.T) {
	one := GetThree(1)
	two := GetThree(2)

	if one.CanMerge(two) != true || two.CanMerge(one) != true {
		t.Error("Between threes that are 1 and 2 should be able to merge to 3")
	}

	greaterThanThree := GetThree(3)
	six := GetThree(4)

	if greaterThanThree.CanMerge(six) == true {
		t.Error("Between threes that are greater than 3 should be able to merge only equal value self")
	}

	if greaterThanThree.CanMerge(greaterThanThree) != true {
		t.Error("Between threes that are greater than 3 should be able to merge only equal value self")
	}
}

func TestShouldMergeTwoThreeToOne(t *testing.T) {
	one := GetThree(1)
	two := GetThree(2)

	if one.Merge(two) != GetThree(3) {
		t.Error("To merge 1 and 2 should be 3")
	}
}
