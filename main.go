package main

import (
	//	"fmt"
	t "github.com/coporter/rogue-golang/terminal"
	//	"github.com/nsf/termbox-go"
)

func main() {
	t.Open()
	defer t.Close()

loop:
	for {
		key := t.GetKeyPress()
		if key == t.KeyEscape {
			break loop
		}
		t.PrintRune(rune(key), 0, 0, t.ColorRed, t.ColorDefault)
		t.Flush()
	}
}

//
//type Thing struct {
//	x, y             int
//	symbol           rune
//	fgColor, bgColor termbox.Attribute
//}
//
//func (thing Thing) draw() {
//	termbox.SetCell(thing.x, thing.y, thing.symbol, thing.fgColor, thing.bgColor)
//}
//
