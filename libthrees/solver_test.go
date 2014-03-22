package libthrees

import (
	"testing"
)

func TestCanMakeEmptyBasement(t *testing.T) {
	basement := GetBasement()

	for _, v := range(basement.area) {
		for _, v := range(v) {
			if !v.IsEmpty() {
				t.Error("Empty basement have to have only nil as threes")
			}
		}
	}
}

func TestCanMakeInitializedBasementWithThrees(t *testing.T) {
	accessor := []ValueDefiner{&PosValue{0,0,GetThree(1)}}
	basement,_ := GetInitializedBasement(accessor)

	if basement.area[0][0] != GetThree(1) {
		t.Errorf("Initialized basement has some threes in given basements.")
	}

	accessor = []ValueDefiner{&PosValue{-1,0,GetThree(1)}}
	if _, error := GetInitializedBasement(accessor); error == nil {
		t.Errorf("GetInitializedBasement should be return error with negative position")
	}

	accessor = []ValueDefiner{&PosValue{0,-1,GetThree(1)}}
	if _, error := GetInitializedBasement(accessor); error == nil {
		t.Errorf("GetInitializedBasement should be return error with negative position")
	}
}

func TestCanGetAtThePosition(t *testing.T) {
	positions := []Pos {
		Pos{-1,0},
		Pos{-1, -1},
		Pos{0, -1}}
		
	basement := GetBasement()

	for _, v := range(positions) {
		if _, err := basement.At(v); err == nil {
			t.Errorf("Position %s should not be able to get some value", v)
		}
	}
}

func TestCanSolveIfItHaveMerge(t *testing.T) {

	accessor := []ValueDefiner{
		&PosValue{0,0, GetThree(1)},
		&PosValue{0,1, GetThree(2)}}

	basement, _ := GetInitializedBasement(accessor)

	basement, err := basement.Solve(UP)
	if err != nil {
		t.Error("Should be able to solve this to UP direction that has 1 and 2")
	}

	if v, _ := basement.At(Pos{0, 0}); v != GetThree(3) {
		t.Errorf("Merged threes with 1 and 2 should be 3, but %s", v)
	}
}

func TestCanSeparateEachLineOfDirection(t *testing.T) {
	
}

















