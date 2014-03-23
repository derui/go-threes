package main

import (
	"github.com/derui/goansi"
	"github.com/derui/go-threes/libthrees"
	"os/exec"
	"os"
	"math/rand"
	"fmt"
	"strings"
)

const (
	CARD_WIDTH = 7
	CARD_HEIGHT = 3

	CANVAS_SIZE = CARD_WIDTH * libthrees.BASEMENT_SIZE + libthrees.BASEMENT_SIZE + 1
)

var (
	NUM_RANGE_1 = 4
	NUM_RANGE_2 = 4
	NUM_RANGE_3 = 4
	NUM_RANGE_CARD_INITIALIZE = []int{6, 8}
)

var nextCard Card

func initialize() {
	cmd := exec.Command("stty", "-F", "/dev/tty", "-echo", "cbreak", "min", "1")
	cmd.Run()

	goansi.HideCursor()

	updateNextCard()
}

func finalize() {
	goansi.ShowCursor();
	goansi.Erase()
}

func initializeBasement() (libthrees.Basement) {
	cards := []*Card{}
	const size = libthrees.BASEMENT_SIZE

	numCards := NUM_RANGE_CARD_INITIALIZE[rand.Intn(NUM_RANGE_CARD_INITIALIZE[1] - NUM_RANGE_CARD_INITIALIZE[0])]
	num3 := rand.Intn(NUM_RANGE_3) + 1
	num1 := rand.Intn(numCards - num3) + 1
	num2 := numCards - num3 - num1

	cards = append(cards, createCards(num1, 1)...)
	cards = append(cards, createCards(num2, 2)...)
	cards = append(cards, createCards(num3, 3)...)
	cards = append(cards, createCards(size * size - numCards, 0)...)

	permIndex := rand.Perm(libthrees.BASEMENT_SIZE * libthrees.BASEMENT_SIZE)
	positions := GetAllCombPos()

	for index, v := range cards {
		cards[index] = &Card{v.target, positions[permIndex[index]]}
	}

	param := make([]libthrees.ValueDefiner, len(cards))
	for index, v := range cards {
		param[index] = v
	}

	basement,_ := libthrees.GetInitializedBasement(param)
	return basement
}

func updateNextCard() {
	nextCard.target = libthrees.GetThree(rand.Intn(3) + 1)
}

func createCards(num int, numOfThree int) []*Card {

	result := []*Card{}

	for i := 0;i < num;i++ {
		result = append(result, &Card{libthrees.GetThree(numOfThree), libthrees.Pos{}})
	}

	return result
}

func renderBasement(basement libthrees.Basement) {
	goansi.Erase()

	// 次の項目を表示するための領域を作成する
	Render(&nextCard, CANVAS_SIZE / 2 - CARD_WIDTH / 2, 1)

	baseY := CARD_HEIGHT + 4
	goansi.MoveTo(0, baseY)

	area := basement.GetArea()

	for _, card := range area {
		x := card.X()
		y := card.Y()

		Render(card,x * (CARD_WIDTH + 1) + 1, y * (CARD_HEIGHT + 1) + 1 + baseY)
	}
}

// 全体のループ処理を行う。ハンドラがfalseを返した時点でループから抜ける
func loop(onRender func(), onLoop func(byte) bool) {

	input := make([]byte, 1)

	for {
		onRender()
		os.Stdin.Read(input)

		if !onLoop(input[0]) {
			break
		}
	}
}

func makeSolver(generator func () libthrees.Three) libthrees.Solver {

	return func(indices []int) (int, libthrees.Three) {
		
		if len(indices) == 1 {
			return indices[0], generator()
		}

		target := rand.Intn(len(indices))
		return indices[target], generator()
	}
}

func calcScore(t libthrees.Basement) int {
	area := t.GetArea()
	score := 0

	for _, v := range area {
		score += v.Value().Score()
	}

	return score
}

func main() {

	initialize()

	basement := initializeBasement()

	generator := func () libthrees.Three {
		return nextCard.Value()
	}

	loop(
		func() {
			renderBasement(basement)
		},
		func(c byte) bool {

			if !basement.CanSomeSolve() {
				return false
			}
			switch c {
			case 'h':
				if basement.Solve(libthrees.LEFT, makeSolver(generator)) {
					updateNextCard()
				}
				return true
			case 'j':
				if basement.Solve(libthrees.DOWN, makeSolver(generator)) {
					updateNextCard()
				}
				return true
			case 'k':
				if basement.Solve(libthrees.UP, makeSolver(generator)) {
					updateNextCard()
				}
				return true
			case 'l':
				if basement.Solve(libthrees.RIGHT, makeSolver(generator)) {
					updateNextCard()
				}
				return true
			case 'q':
				return false
			default:
				return true
			}
		})

	showScore(basement)

	finalize()
}

func showScore(t libthrees.Basement) {

	goansi.Erase()

	goansi.MoveTo(0, 0)

	fmt.Print("+" + strings.Repeat("-", CANVAS_SIZE - 2) + "+")

	goansi.MoveTo(0, 1)
	whitespace := goansi.OnWhite(strings.Repeat(" ", CANVAS_SIZE - 2))
	fmt.Print("|" + whitespace + "|")

	goansi.MoveTo(0, 2)
	fmt.Print("+" + strings.Repeat("-", CANVAS_SIZE - 2) + "+")

	goansi.MoveTo(3, 1)
	score := calcScore(t)

	fmt.Print(goansi.OnWhite(goansi.Black(fmt.Sprintf("Score ==> %d", score))))

	input := make([]byte, 1)
	os.Stdin.Read(input)
}

















