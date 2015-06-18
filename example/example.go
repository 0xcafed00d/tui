package main

import (
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/go_sandbox/geom"
	"github.com/simulatedsimian/tui"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	doQuit := false

	//logDisp := tui.ScrollingTextOutput{1, 20, 80, 10, nil}
	cmdInput := tui.MakeTextInputField(10, 18, func(cmd string) {
		//logDisp.WriteLine(cmd)
		if cmd == "q" {
			doQuit = true
		}
	})

	dl := tui.DisplayList{}
	dl.AddElement(cmdInput)
	//	dl.AddElement(&logDisp)
	dl.AddElement(tui.MakeStaticText(geom.Rectangle{geom.Coord{0, 0}, geom.Coord{0, 0}}, "StaticText"))

	dl.Draw()
	termbox.Flush()

	for !doQuit {
		ev := termbox.PollEvent()

		if ev.Type == termbox.EventKey {
			dl.HandleInput(ev.Key, ev.Ch)
			dl.Draw()
			termbox.Flush()
		}

		if ev.Type == termbox.EventResize {
			termbox.Flush()
		}
	}

}
