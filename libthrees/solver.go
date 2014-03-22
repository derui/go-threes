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

// The type of the basement for Threes.
// This type contains all parameters of Threes to solve it.
type Basement struct {

	// Area for solving threes.
	area [BASEMENT_SIZE][BASEMENT_SIZE]three

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

// 指定された方向について解き、新しい結果となるBasementを返す。
// 指定された方向について解けなかった場合は、errorを返す。
func (t Basement) Solve(dir Direction) (Basement, error) {

	switch dir {
	case UP:
	case DOWN:
	case LEFT:
	case RIGHT:
	}
	return GetBasement(), nil
}




















