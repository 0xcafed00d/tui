package tui

import (
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/go_sandbox/geom"
)

type StaticText struct {
	TUIElement
	Text string
}

func MakeStaticText(pos geom.Rectangle, text string) *StaticText {
	return &StaticText{TUIElement{pos}, text}
}

func (t *StaticText) HandleInput(k termbox.Key, r rune) {
}

func (t *StaticText) Draw() {
	printAtDef(t.Min.X, t.Min.Y, t.Text)
}

func (t *StaticText) GiveFocus() bool {
	return false
}
