package tui

import (
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/go_sandbox/geom"
)

func printAt(x, y int, s string, fg, bg termbox.Attribute) {
	for _, r := range s {
		termbox.SetCell(x, y, r, fg, bg)
		x++
	}
}

func printAtDef(x, y int, s string) {
	printAt(x, y, s, termbox.ColorDefault, termbox.ColorDefault)
}

func clearRect(rect geom.Rectangle, c rune, fg, bg termbox.Attribute) {
	w, h := termbox.Size()
	sz := geom.RectangleFromSize(geom.Coord{w, h})

	toClear, ok := geom.RectangleIntersection(rect, sz)
	if ok {
		for y := toClear.Min.Y; y < toClear.Max.Y; y++ {
			for x := toClear.Min.X; x < toClear.Max.X; x++ {
				termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
			}
		}
	}
}

func clearRectDef(rect geom.Rectangle) {
	clearRect(rect, '.', termbox.ColorDefault, termbox.ColorDefault)
}

type DisplayElement interface {
	GiveFocus() bool
	HandleInput(k termbox.Key, r rune)
	Draw()
}

type DisplayList struct {
	list       []DisplayElement
	focusIndex int
}

func (dl *DisplayList) AddElement(elem DisplayElement) {
	dl.list = append(dl.list, elem)
}

func (dl *DisplayList) Draw() {
	w, h := termbox.Size()
	clearRectDef(geom.RectangleFromSize(geom.Coord{w, h}))

	for _, elem := range dl.list {
		elem.Draw()
	}
}

func (dl *DisplayList) NextFocus() {
	if dl.list != nil && len(dl.list) > 0 {
		for {
			dl.focusIndex++
			if dl.focusIndex >= len(dl.list) {
				dl.focusIndex = 0
			}

			if dl.list[dl.focusIndex].GiveFocus() {
				break
			}
		}
	}
}

func (dl *DisplayList) PrevFocus() {
	if dl.list != nil && len(dl.list) > 0 {
		for {
			dl.focusIndex--
			if dl.focusIndex < 0 {
				dl.focusIndex = len(dl.list)
			}

			if dl.list[dl.focusIndex].GiveFocus() {
				break
			}
		}
	}
}

func (dl *DisplayList) HandleInput(k termbox.Key, r rune) {

	if dl.list != nil && len(dl.list) > 0 {
		if k == termbox.KeyTab {
			dl.NextFocus()
		} else {
			dl.list[dl.focusIndex].HandleInput(k, r)
		}
	}
}

type TUIElement struct {
	geom.Rectangle
}
