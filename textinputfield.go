package tui

import (
	"container/list"
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/runes"
	"reflect"
	"unicode"
)

type InputHandler func(inp string)

type TextInputField struct {
	x, y       int
	inp        []rune
	cursorLoc  int
	inpHandler InputHandler
	hasFocus   bool
	history    *list.List
	histPos    *list.Element
}

func MakeTextInputField(x, y int, inpHandler InputHandler) *TextInputField {
	return &TextInputField{x, y, nil, 0, inpHandler, false, list.New(), nil}
}

func (t *TextInputField) HandleInput(k termbox.Key, r rune) {
	if k == termbox.KeyEnter {
		t.inp = runes.Trim(t.inp, unicode.IsSpace)
		if len(t.inp) > 0 {
			if t.histPos != nil && reflect.DeepEqual(t.histPos.Value, t.inp) {
				t.history.MoveToBack(t.histPos)
			} else {
				t.history.PushBack(t.inp)
			}
			t.histPos = nil

			t.inpHandler(string(t.inp))
			t.inp = nil
			t.cursorLoc = 0
		}
	}

	if k == termbox.KeyEsc {
		t.inp = nil
		t.cursorLoc = 0
	}

	if k == termbox.KeyArrowLeft {
		if t.cursorLoc > 0 {
			t.cursorLoc--
		}
	}

	if k == termbox.KeyArrowRight {
		if t.cursorLoc < len(t.inp) {
			t.cursorLoc++
		}
	}

	if k == termbox.KeyArrowUp {
		if t.history.Len() > 0 {
			if t.histPos == nil {
				t.histPos = t.history.Back()
			} else {
				t.histPos = t.histPos.Prev()
				if t.histPos == nil {
					t.histPos = t.history.Front()
				}
			}
			t.inp = runes.CloneSlice(t.histPos.Value.([]rune))
			t.cursorLoc = len(t.inp)
		}
	}

	if k == termbox.KeyArrowDown {
		if t.history.Len() > 0 {
			if t.histPos != nil {
				t.histPos = t.histPos.Next()
			}

			t.inp = nil
			if t.histPos != nil {
				t.inp = append(t.inp, t.histPos.Value.([]rune)...)
			}
			t.cursorLoc = len(t.inp)
		}
	}

	if r > ' ' {
		t.inp = runes.InsertAt(t.inp, r, t.cursorLoc)
		t.cursorLoc++
	}

	if k == 32 {
		t.inp = runes.InsertAt(t.inp, ' ', t.cursorLoc)
		t.cursorLoc++
	}

	if t.cursorLoc > 0 && len(t.inp) > 0 && (k == termbox.KeyBackspace || k == termbox.KeyBackspace2) {
		t.inp = runes.DeleteAt(t.inp, t.cursorLoc-1)
		t.cursorLoc--
	}

	if k == termbox.KeyDelete && t.cursorLoc < len(t.inp) {
		t.inp = runes.DeleteAt(t.inp, t.cursorLoc)
	}

	if k == termbox.KeyHome {
		t.cursorLoc = 0
	}

	if k == termbox.KeyEnd {
		t.cursorLoc = len(t.inp)
	}

	termbox.SetCursor(t.x+t.cursorLoc, t.y)
	//	printAtDef(t.x, t.y+1, fmt.Sprintf("%v, %v               ", k, r))
}

func (t *TextInputField) Draw() {
	printAtDef(t.x, t.y, string(t.inp)+" ")
}

func (t *TextInputField) GiveFocus() bool {
	termbox.SetCursor(t.x+t.cursorLoc, t.y)
	return true
}
