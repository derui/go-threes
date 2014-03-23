package main

import (
	"github.com/derui/go-threes/libthrees"
	"github.com/derui/goansi"
	"fmt"
	"strings"
)

type Card struct {
	target libthrees.Three
	pos libthrees.Pos
}

func (t *Card) X() int {
	return t.pos.X()
}

func (t *Card) Y() int {
	return t.pos.Y()
}

func (t *Card) Value() libthrees.Three {
	return t.target
}

func Render(t libthrees.ValueDefiner, x,y int) {

	three := t.Value().String()
	goansi.MoveTo(x, y)
	
	fmt.Print("+" + strings.Repeat("-", CARD_WIDTH) + "+")
	for i := 0; i < CARD_HEIGHT; i++ {
		goansi.MoveTo(x, y + i + 1)


		whitespace := goansi.OnWhite(strings.Repeat(" ", CARD_WIDTH))
		if t.Value().IsEmpty() {
			whitespace = goansi.OnBlack(strings.Repeat(" ", CARD_WIDTH))
		}
		fmt.Print("|" + whitespace + "|")
	}
	goansi.MoveTo(x, y + CARD_HEIGHT + 1)
	fmt.Print("+" + strings.Repeat("-", CARD_WIDTH) + "+")

	goansi.MoveTo(x + 1 + CARD_WIDTH / 2 - (len(three) - 1), y + CARD_HEIGHT / 2 + 1)

	switch t.Value().Value() {
	case 0:
	case 1:
		fmt.Print(goansi.Blue(goansi.OnWhite(three)))
	case 2:
		fmt.Print(goansi.Red(goansi.OnWhite(three)))
	default:
		fmt.Print(goansi.Black(goansi.OnWhite(three)))
	}
}



















