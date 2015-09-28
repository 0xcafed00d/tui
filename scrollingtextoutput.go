package tui

import (
	"github.com/simulatedsimian/rect"
)

type ScrollingTextOutput struct {
	TUIElement
	text []string
}

func MakeScrollingTextOutput(pos rect.Rectangle) *ScrollingTextOutput {
	return &ScrollingTextOutput{TUIElement{pos}, nil}
}

func (t *ScrollingTextOutput) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (t *ScrollingTextOutput) WriteLine(l string) {
	t.text = append(t.text, l)
}

func (t *ScrollingTextOutput) Draw() {
	clearRectDef(t.TUIElement.Rectangle)

	start := 0

	if len(t.text) > t.Height() {
		start = len(t.text) - t.Height()
	}

	y := t.Min.Y
	for l := start; l < len(t.text); l++ {
		printAtDef(t.Min.X, y, t.text[l])
		y++
	}
}
