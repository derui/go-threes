package libthrees

import (
	"fmt"
)

// Basementへのアクセッサ。設定や取得などは、すべてこのアクセッサを経由して行う
type Accessor interface {
	X() int
	Y() int
}

// 値を含めたアクセッサ
type ValueDefiner interface {
	Accessor
	Value() Three
}

// ある位置を表すための構造体
type Pos struct {
	x int
	y int
}

func (t *Pos) X() int {
	return t.x
}

func (t *Pos) Y() int {
	return t.y
}

func (t *Pos) String() string {
	return fmt.Sprintf("(%d,%d)", t.x, t.y)
}

func GetPos(x, y int) Pos {
	return Pos{x, y}
}

// ある位置と値をまとめて表現するための構造体
type PosValue struct {
	x     int
	y     int
	value Three
}

func (t *PosValue) X() int {
	return t.x
}

func (t *PosValue) Y() int {
	return t.y
}

func (t *PosValue) Value() Three {
	return t.value
}
