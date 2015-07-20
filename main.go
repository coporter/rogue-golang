package main

import (
	//	"fmt"
	"github.com/coporter/rogue-golang/terminal"
	//	"github.com/nsf/termbox-go"
	"time"
)

func main() {
	//err := termbox.Init()
	//if err != nil {
	//	panic(err)
	//}
	//defer termbox.Close()

	terminal.Init()
	terminal.PrintRune('@', 5, 5)
	terminal.PrintLine("zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz", 5, 6)
	time.Sleep(time.Second * 2)
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
//func printLine(str string, x int, y int) {
//	for i := range str {
//		//		termbox.SetCell(x+i, y, rune(str[i]), termbox.ColorDefault, termbox.ColorDefault)
//		printRune(rune(str[i]), x+i, y)
//	}
//}
//
//func printRune(r rune, x int, y int) {
//	termbox.SetCell(x, y, rune(r), termbox.ColorDefault, termbox.ColorDefault)
//}
//
//func main() {
//	err := termbox.Init()
//	if err != nil {
//		panic(err)
//	}
//	defer termbox.Close()
//
//	player := Thing{0, 0, '@', termbox.ColorDefault, termbox.ColorRed}
//
//loop:
//	for {
//		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
//		switch ev := termbox.PollEvent(); ev.Type {
//		case termbox.EventKey:
//			switch ev.Ch {
//			case 'k':
//				printLine("UP!", 0, 0)
//				player.y -= 1
//			case 'h':
//				printLine("LEFT!", 0, 0)
//				player.x -= 1
//			case 'j':
//				printLine("DOWN!", 0, 0)
//				player.y += 1
//			case 'l':
//				printLine("RIGHT!", 0, 0)
//				player.x += 1
//			case 'q':
//				break loop
//				//default:
//				//printLine("Invalid command!", 0, 0)
//			}
//			switch ev.Key {
//			case termbox.KeyArrowUp:
//				printLine("UP!", 0, 0)
//				player.y -= 1
//			case termbox.KeyArrowLeft:
//				printLine("LEFT!", 0, 0)
//				player.x -= 1
//			case termbox.KeyArrowDown:
//				printLine("DOWN!", 0, 0)
//				player.y += 1
//			case termbox.KeyArrowRight:
//				printLine("RIGHT!", 0, 0)
//				player.x += 1
//			case termbox.KeyEsc:
//				break loop
//				//default:
//				//printLine("Invalid command!", 0, 0)
//			}
//			printLine(fmt.Sprintf("EventKey: k: %d, c: %c, mod: %d", ev.Key, ev.Ch, ev.Mod), 0, 2)
//		}
//		player.draw()
//		termbox.SetCell(10, 10, 'Z', color.Red, color.Default)
//		termbox.Flush()
//	}
//}
