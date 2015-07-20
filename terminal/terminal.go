package terminal

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

type Color uint16

const (
	ColorDefault = termbox.ColorDefault
	ColorBlack   = termbox.ColorBlack
	ColorRed     = termbox.ColorRed
	ColorGreen   = termbox.ColorGreen
	ColorYellow  = termbox.ColorYellow
	ColorBlue    = termbox.ColorBlue
	ColorMagenta = termbox.ColorMagenta
	ColorCyan    = termbox.ColorCyan
	ColorWhite   = termbox.ColorWhite
)

func Init() {
	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAA")
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
}

func PrintRune(r rune, x int, y int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}

func PrintLine(str string, x int, y int) {
	for i := range str {
		PrintRune(rune(str[i]), x+i, y)
	}
}
