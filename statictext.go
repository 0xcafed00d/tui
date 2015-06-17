package tui

import (
	"github.com/nsf/termbox-go"
)

type StaticText struct {
	x, y int
	text string
}

func (t *StaticText) HandleInput(k termbox.Key, r rune) {
}

func (t *StaticText) Draw() {
	printAtDef(t.x, t.y, t.text)
}

func (t *StaticText) GiveFocus() bool {
	return false
}
