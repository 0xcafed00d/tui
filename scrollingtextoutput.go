package tui

import (
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/go_sandbox/geom"
)

type ScrollingTextOutput struct {
	x, y int
	w, h int
	text []string
}

func (t *ScrollingTextOutput) HandleInput(k termbox.Key, r rune) {
}

func (t *ScrollingTextOutput) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (t *ScrollingTextOutput) WriteLine(l string) {
	t.text = append(t.text, l)
}

func (t *ScrollingTextOutput) Draw() {
	clearRectDef(geom.RectangleFromPosSize(geom.Coord{t.x, t.y}, geom.Coord{t.w, t.h}))

	start := 0

	if len(t.text) > t.h {
		start = len(t.text) - t.h
	}

	y := t.y
	for l := start; l < len(t.text); l++ {
		printAtDef(t.x, y, t.text[l])
		y++
	}
}

func (t *ScrollingTextOutput) GiveFocus() bool {
	return false
}
