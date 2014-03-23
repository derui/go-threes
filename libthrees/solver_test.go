package libthrees

import (
	"testing"
)

func TestCanMakeEmptyBasement(t *testing.T) {
	basement := GetBasement()

	for _, v := range basement.area {
		for _, v := range v {
			if !v.IsEmpty() {
				t.Error("Empty basement have to have only nil as threes")
			}
		}
	}
}

func TestCanMakeInitializedBasementWithThrees(t *testing.T) {
	accessor := []ValueDefiner{&PosValue{0, 0, GetThree(1)}}
	basement, _ := GetInitializedBasement(accessor)

	if basement.area[0][0] != GetThree(1) {
		t.Errorf("Initialized basement has some threes in given basements.")
	}

	accessor = []ValueDefiner{&PosValue{-1, 0, GetThree(1)}}
	if _, error := GetInitializedBasement(accessor); error == nil {
		t.Errorf("GetInitializedBasement should be return error with negative position")
	}

	accessor = []ValueDefiner{&PosValue{0, -1, GetThree(1)}}
	if _, error := GetInitializedBasement(accessor); error == nil {
		t.Errorf("GetInitializedBasement should be return error with negative position")
	}
}

func TestCanGetAtThePosition(t *testing.T) {
	positions := []Pos{
		Pos{-1, 0},
		Pos{-1, -1},
		Pos{0, -1}}

	basement := GetBasement()

	for _, v := range positions {
		if _, err := basement.At(v); err == nil {
			t.Errorf("Position %s should not be able to get some value", v)
		}
	}
}

func TestShouldGetErrorIfNotMerged(t *testing.T) {
	base := []ValueDefiner{}

	for x := 0; x < BASEMENT_SIZE; x++ {
		for y := 0; y < BASEMENT_SIZE; y++ {
			base = append(base, &PosValue{x, y, GetThree(1)})
		}
	}

	basement, _ := GetInitializedBasement(base)

	for _, d := range []Direction{UP, DOWN, LEFT, RIGHT} {
		if basement.Solve(d, func (indices []int) (int, Three) {return 0, GetThree(1)}) {
			t.Errorf("Should return some error if can not any solving : %s", d)
		}
	}
}

func TestCanSolveIfItHaveMerge(t *testing.T) {

	accessor := []ValueDefiner{
		&PosValue{0, 0, GetThree(1)},
		&PosValue{0, 1, GetThree(2)}}

	basement, _ := GetInitializedBasement(accessor)

	solved := basement.Solve(UP, func (indices []int) (int, Three) {
		return 0, GetThree(1)
	})
	if !solved {
		t.Error("Should be able to solve this to UP direction that has 1 and 2")
	}

	if v, _ := basement.At(Pos{0, 0}); v != GetThree(3) {
		t.Errorf("Merged threes with 1 and 2 should be 3, but %s", v)
	}

	if v, _ := basement.At(Pos{0, BASEMENT_SIZE - 1}); v != GetThree(1) {
		t.Errorf("Merged line should be appended a new three from newline")
	}
}

func TestCanDetectToBeAbleToSolve(t *testing.T) {
	accessor := []ValueDefiner{}
	for y := 0; y < BASEMENT_SIZE; y++ {
		for x := 0; x < BASEMENT_SIZE; x++ {
			accessor = append(accessor, &PosValue{x, y, GetThree(1)})
		}
	}

	basement, _ := GetInitializedBasement(accessor)

	if basement.CanSolve(UP) {
		t.Error("Filled with only 1 is not able to solve for UP, but say can solve")
	}

	if basement.CanSolve(LEFT) {
		t.Error("Filled with only 1 is not able to solve for LEFT, but say can solve")
	}

	if basement.CanSolve(RIGHT) {
		t.Error("Filled with only 1 is not able to solve for RIGHT, but say can solve")
	}

	if basement.CanSolve(DOWN) {
		t.Error("Filled with only 1 is not able to solve DOWN, but say can solve")
	}

	if basement.CanSomeSolve() {
		t.Error("Filled with only 1 is not able to solve for some direction, but say can solve")
	}
}

func TestCanSolveIfOverrideEmpty(t *testing.T) {

	accessor := []ValueDefiner{
		&PosValue{0, 1, GetThree(1)},
		&PosValue{0, 2, GetThree(2)}}

	basement, _ := GetInitializedBasement(accessor)

	solved := basement.Solve(UP, func (indices []int) (int, Three) {
		return 0, GetThree(1)
	})
	if !solved {
		t.Error("Should be able to solve this to UP direction that has 1 and 2")
	}

	if v, _ := basement.At(Pos{0, 0}); v != GetThree(1) {
		t.Errorf("Override empty if own or other is empty")
	}

	if v, _ := basement.At(Pos{0, BASEMENT_SIZE - 1}); v != GetThree(1) {
		t.Errorf("Merged line should be appended a new three from newline")
	}
}


var positionsMap = map[Direction][][]Pos{
	UP: [][]Pos{
		[]Pos{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		[]Pos{{1, 0}, {1, 1}, {1, 2}, {1, 3}},
		[]Pos{{2, 0}, {2, 1}, {2, 2}, {2, 3}},
		[]Pos{{3, 0}, {3, 1}, {3, 2}, {3, 3}}},

	DOWN: [][]Pos{
		[]Pos{{0, 3}, {0, 2}, {0, 1}, {0, 0}},
		[]Pos{{1, 3}, {1, 2}, {1, 1}, {1, 0}},
		[]Pos{{2, 3}, {2, 2}, {2, 1}, {2, 0}},
		[]Pos{{3, 3}, {3, 2}, {3, 1}, {3, 0}}},

	LEFT: [][]Pos{
		[]Pos{{0, 0}, {1, 0}, {2, 0}, {3, 0}},
		[]Pos{{0, 1}, {1, 1}, {2, 1}, {3, 1}},
		[]Pos{{0, 2}, {1, 2}, {2, 2}, {3, 2}},
		[]Pos{{0, 3}, {1, 3}, {2, 3}, {3, 3}}},

	RIGHT: [][]Pos{
		[]Pos{{3, 0}, {2, 0}, {1, 0}, {0, 0}},
		[]Pos{{3, 1}, {2, 1}, {1, 1}, {0, 1}},
		[]Pos{{3, 2}, {2, 2}, {1, 2}, {0, 2}},
		[]Pos{{3, 3}, {2, 3}, {1, 3}, {0, 3}}},
}

func TestShouldGetPositionsForSeparations(t *testing.T) {
	for _, k := range []Direction{UP, DOWN, LEFT, RIGHT} {
		expect := positionsMap[k]

		actual := getPositions(k)

		for yi, l := range expect {
			for xi, v := range l {

				if v != actual[yi][xi] {
					t.Error("To get positions for %s at (%d, %d) should be %s, but %s",
						k, xi, yi, v, actual[yi][xi])
				}
			}
		}
	}
}

func TestCanSeparateEachLineOfDirection(t *testing.T) {
	firstLine := []ValueDefiner{
		&PosValue{0, 0, GetThree(1)},
		&PosValue{0, 1, GetThree(2)},
		&PosValue{0, 2, GetThree(3)},
		&PosValue{0, 3, GetThree(1)}}

	base := append(firstLine, []ValueDefiner{
		&PosValue{1, 0, GetThree(1)},
		&PosValue{2, 0, GetThree(2)},
		&PosValue{3, 0, GetThree(3)}}...)

	basement, _ := GetInitializedBasement(base)

	separated := basement.separateWithDirection(UP)

	if len(separated) != BASEMENT_SIZE {
		t.Error("Separation with direction are number of separated as basement_size")
	}

	first := separated[0]
	for i, v := range first.data {
		if firstLine[i].Value() != v {
			t.Errorf("A value on the first line at %d should be expected %s, but get %s",
				i, base[i].Value(), v)
		}
	}
}
