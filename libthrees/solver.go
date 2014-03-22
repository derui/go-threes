package libthrees

import (
	"fmt"
)

// Define solver for Threes and Provide base types for solving.

const (
	BASEMENT_SIZE = 4
)

type Direction int

const (
	UP Direction = iota
	DOWN Direction = iota
	LEFT Direction = iota
	RIGHT Direction = iota
)

func (t Direction) String() string {
	switch t {
	case UP:
		return "UP"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	case RIGHT:
		return "RIGHT"
	default:
		return ""
	}
}

// The type of the basement for Threes.
// This type contains all parameters of Threes to solve it.
type Basement struct {

	// Area for solving threes.
	area [BASEMENT_SIZE][BASEMENT_SIZE]three
}

// 元となるbasementから、指定されたDirectionで取得された列を表す
type Line struct {
	data []three

	dir Direction
}

func GetBasement() Basement {
	return Basement{[BASEMENT_SIZE][BASEMENT_SIZE]three {
		{empty(), empty(), empty(), empty()},
		{empty(), empty(), empty(), empty()},
		{empty(), empty(), empty(), empty()},
		{empty(), empty(), empty(), empty()}}}
}

func between(v int) bool {
	return 0 <= v && v < BASEMENT_SIZE
}

// 初期化されたベースを返す。
func GetInitializedBasement(accessors []ValueDefiner) (Basement, error) {

	basement := GetBasement();

	for _, v := range(accessors) {
		if !between(v.X()) || !between(v.Y()) {
			return GetBasement(), fmt.Errorf("the position of basement as x and y must 0 < x < BASEMENT_SIZE")
		}

		basement.area[v.Y()][v.X()] = v.Value()
	}

	return basement, nil
}

// 指定された位置の値を取得する
func (t Basement) At(pos Pos) (three, error) {

	if !between(pos.X()) || !between(pos.Y()) {
		return empty(), fmt.Errorf("(x, y) = (%d, %d) is not contains from basement", pos.X(), pos.Y())
	}

	return t.area[pos.Y()][pos.X()], nil
}

func getPositions(dir Direction) [][]Pos {
	switch dir {
	case UP:
		ret := make([][]Pos, 0)
		for xi := 0; xi < BASEMENT_SIZE; xi++ {
			line := make([]Pos, 0)
			for yi := 0; yi < BASEMENT_SIZE; yi++ {
				line = append(line, Pos{xi, yi})
			}
			ret = append(ret, line)
		}

		return ret
	case DOWN:
		ret := make([][]Pos, 0)
		for xi := 0; xi < BASEMENT_SIZE; xi++ {
			line := make([]Pos, 0)
			for yi := BASEMENT_SIZE - 1; yi >= 0 ; yi-- {
				line = append(line, Pos{xi, yi})
			}
			ret = append(ret, line)
		}

		return ret
	case LEFT:
		ret := make([][]Pos, 0)
		for yi := 0; yi < BASEMENT_SIZE; yi++ {
			line := make([]Pos, 0)
			for xi := 0; xi < BASEMENT_SIZE ; xi++ {
				line = append(line, Pos{xi, yi})
			}
			ret = append(ret, line)
		}

		return ret
	case RIGHT:
		ret := make([][]Pos, 0)
		for yi := 0; yi < BASEMENT_SIZE; yi++ {
			line := make([]Pos, 0)
			for xi := BASEMENT_SIZE - 1; xi >= 0; xi-- {
				line = append(line, Pos{xi, yi})
			}
			ret = append(ret, line)
		}

		return ret
	default:
		panic("Unknown Direction value!")
	}
}

func (t Basement) separateWithDirection(dir Direction) []Line {

	positions := getPositions(dir)

	ret := []Line{}
	for _, linePositions := range positions {

		line := Line{}
		
		for _, pos := range linePositions {
			v, _ := t.At(pos)
			line.data = append(line.data, v)
		}
		ret = append(ret, line)
	}

	return ret
}

func MakeNewLine(pos int, newThree three) [BASEMENT_SIZE]three {
	if !between(pos) {
		panic("`pos' should between 0 and BASEMENT_SIZE")
	}

	ret := [BASEMENT_SIZE]three{empty(), empty(), empty(), empty()}
	ret[pos] = newThree

	return ret
}

// 指定された方向について解き、新しい結果となるBasementを返す。
// 指定された方向について解けなかった場合は、errorを返す。
func (t Basement) Solve(dir Direction, newThrees [BASEMENT_SIZE]three) (Basement, error) {

	lines := t.separateWithDirection(dir)
	// 各行のpositionを元に、元々の場所に再設定する。
	positions := getPositions(dir)

	allError := true
	for index, line := range lines {
		line.data = append(line.data, newThrees[index])
		merged, mergedLine := solveLine(line)

		position := positions[index]

		for mergedLineIndex, v := range mergedLine.data {
			pos := position[mergedLineIndex]
			t.area[pos.Y()][pos.X()] = v
		}

		if merged {
			allError = false
		}
	}

	if allError {
		return t, fmt.Errorf("can't solve")
	}

	return t, nil
}

// ある行について解き、解いた結果を返す。
// ある行自体が解けなかった場合、solvedにfalseが設定される
func solveLine(line Line) (solved bool, ret Line) {

	solveSets := [][]int{{0,1}, {1,2}, {2,3}}
	solved = false
	ret = line
	ret.data = make([]three, len(line.data) - 1)
	copy(ret.data, line.data[1:])

	// 先頭がEmptyの場合は、一個ずれるだけなので、そのまま設定する
	if line.data[0].IsEmpty() {
		return
	}

	for _, set := range solveSets {
		first := set[0]
		second := set[1]

		if line.data[first].CanMerge(line.data[second]) {
			ret.data[first] = line.data[first].Merge(line.data[second])
			ret.data[first + 1] = line.data[second + 1]

			solved = true
		}
	}

	return 
}


















