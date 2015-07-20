package terminal

import (
	//	"fmt"
	"github.com/nsf/termbox-go"
)

type Color uint

const (
	ColorDefault Color = iota
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

func Open() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

func Close() {
	defer termbox.Close()
}

func getTermboxColor(c Color) termbox.Attribute {
	switch c {
	case ColorDefault:
		return termbox.ColorDefault
	case ColorBlack:
		return termbox.ColorBlack
	case ColorRed:
		return termbox.ColorRed
	case ColorGreen:
		return termbox.ColorGreen
	case ColorYellow:
		return termbox.ColorYellow
	case ColorBlue:
		return termbox.ColorBlue
	case ColorMagenta:
		return termbox.ColorMagenta
	case ColorCyan:
		return termbox.ColorCyan
	case ColorWhite:
		return termbox.ColorWhite
	default:
		return termbox.ColorDefault
	}
}

func PrintRune(r rune, x int, y int, fgColor Color, bgColor Color) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(x, y, r, getTermboxColor(fgColor), getTermboxColor(bgColor))
	termbox.Flush()
}

func PrintLine(str string, x int, y int, fgColor Color, bgColor Color) {
	for i := range str {
		PrintRune(rune(str[i]), x+i, y, fgColor, bgColor)
	}
}
