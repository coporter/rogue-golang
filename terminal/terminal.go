package terminal

import (
	"github.com/nsf/termbox-go"
)

// Color for font
type Color uint

// List of all colors
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

// List of keys
const (
	Key0                  rune = '0'
	Key1                       = '1'
	Key2                       = '2'
	Key3                       = '3'
	Key4                       = '4'
	Key5                       = '5'
	Key6                       = '6'
	Key7                       = '7'
	Key8                       = '8'
	Key9                       = '9'
	KeyAmpersand               = '&'
	KeyAt                      = '@'
	KeyBackslash               = '\\'
	KeyBackspace               = '⌫'
	KeyBacktick                = '`'
	KeyCaret                   = '^'
	KeyColon                   = ':'
	KeyComma                   = ','
	KeyDollar                  = '$'
	KeyDoubleQuote             = '"'
	KeyDownArrow               = '↓'
	KeyEquals                  = '='
	KeyEscape                  = '⎋'
	KeyExcelmation             = '!'
	KeyForwardSlash            = '/'
	KeyGreaterThan             = '>'
	KeyLeftArrow               = '←'
	KeyLeftCurlyBracket        = '{'
	KeyLeftParen               = '('
	KeyLeftSquareBracket       = '['
	KeyLessThan                = '<'
	KeyLowerA                  = 'a'
	KeyLowerB                  = 'b'
	KeyLowerC                  = 'c'
	KeyLowerD                  = 'd'
	KeyLowerE                  = 'e'
	KeyLowerF                  = 'f'
	KeyLowerG                  = 'g'
	KeyLowerH                  = 'h'
	KeyLowerI                  = 'i'
	KeyLowerJ                  = 'j'
	KeyLowerK                  = 'k'
	KeyLowerL                  = 'l'
	KeyLowerM                  = 'm'
	KeyLowerN                  = 'n'
	KeyLowerO                  = 'o'
	KeyLowerP                  = 'p'
	KeyLowerQ                  = 'q'
	KeyLowerR                  = 'r'
	KeyLowerS                  = 's'
	KeyLowerT                  = 't'
	KeyLowerU                  = 'u'
	KeyLowerV                  = 'v'
	KeyLowerW                  = 'w'
	KeyLowerX                  = 'x'
	KeyLowerY                  = 'y'
	KeyLowerZ                  = 'z'
	KeyMinus                   = '-'
	KeyPercent                 = '%'
	KeyPeriod                  = '.'
	KeyPipe                    = '|'
	KeyPlus                    = '+'
	KeyPound                   = '#'
	KeyQuestion                = '?'
	KeyReturn                  = '⏎'
	KeyRightArrow              = '→'
	KeyRightCurlyBracket       = '}'
	KeyRightParen              = ')'
	KeyRightSquareBracket      = ']'
	KeySemicolon               = ';'
	KeySingleQuote             = '\''
	KeySpace                   = '␣'
	KeyStar                    = '*'
	KeyTab                     = '↹'
	KeyTilde                   = '~'
	KeyUnderscore              = '_'
	KeyUnknown                 = '␀'
	KeyUpArrow                 = '↑'
	KeyUpperA                  = 'A'
	KeyUpperB                  = 'B'
	KeyUpperC                  = 'C'
	KeyUpperD                  = 'D'
	KeyUpperE                  = 'E'
	KeyUpperF                  = 'F'
	KeyUpperG                  = 'G'
	KeyUpperH                  = 'H'
	KeyUpperI                  = 'I'
	KeyUpperJ                  = 'J'
	KeyUpperK                  = 'K'
	KeyUpperL                  = 'L'
	KeyUpperM                  = 'M'
	KeyUpperN                  = 'N'
	KeyUpperO                  = 'O'
	KeyUpperP                  = 'P'
	KeyUpperQ                  = 'Q'
	KeyUpperR                  = 'R'
	KeyUpperS                  = 'S'
	KeyUpperT                  = 'T'
	KeyUpperU                  = 'U'
	KeyUpperV                  = 'V'
	KeyUpperW                  = 'W'
	KeyUpperX                  = 'X'
	KeyUpperY                  = 'Y'
	KeyUpperZ                  = 'Z'
)

// Open a connection to the terminal
func Open() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
}

// Close the connection to the terminal
func Close() {
	defer termbox.Close()
}

// Get the termbox color given a Color
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

// PrintRune prints a single rune to the screen at (x,y).
// Uses foreground color fgColor, and background color bgColor.
func PrintRune(r rune, x int, y int, fgColor Color, bgColor Color) {
	termbox.SetCell(x, y, r, getTermboxColor(fgColor), getTermboxColor(bgColor))
}

// PrintLine prints a line of text to the screen starting at (x,y).
// Uses foreground color fgColor, and background color bgColor.
func PrintLine(str string, x int, y int, fgColor Color, bgColor Color) {
	for i := range str {
		PrintRune(rune(str[i]), x+i, y, fgColor, bgColor)
	}
}

// Clear the screen.
func Clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

// Flush the buffered output to the screen.
func Flush() {
	termbox.Flush()
}

// GetKeyPress waits for a keypress, then return the key in rune form.
// Returns KeyUnknown if it unable to identify the key.
func GetKeyPress() rune {
	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventKey:
		switch ev.Ch {
		case 'A':
			return KeyUpperA
		case 'B':
			return KeyUpperB
		case 'C':
			return KeyUpperC
		case 'D':
			return KeyUpperD
		case 'E':
			return KeyUpperE
		case 'F':
			return KeyUpperF
		case 'G':
			return KeyUpperG
		case 'H':
			return KeyUpperH
		case 'I':
			return KeyUpperI
		case 'J':
			return KeyUpperJ
		case 'K':
			return KeyUpperK
		case 'L':
			return KeyUpperL
		case 'M':
			return KeyUpperM
		case 'N':
			return KeyUpperN
		case 'O':
			return KeyUpperO
		case 'P':
			return KeyUpperP
		case 'Q':
			return KeyUpperQ
		case 'R':
			return KeyUpperR
		case 'S':
			return KeyUpperS
		case 'T':
			return KeyUpperT
		case 'U':
			return KeyUpperU
		case 'V':
			return KeyUpperV
		case 'W':
			return KeyUpperW
		case 'X':
			return KeyUpperX
		case 'Y':
			return KeyUpperY
		case 'Z':
			return KeyUpperZ
		case 'a':
			return KeyLowerA
		case 'b':
			return KeyLowerB
		case 'c':
			return KeyLowerC
		case 'd':
			return KeyLowerD
		case 'e':
			return KeyLowerE
		case 'f':
			return KeyLowerF
		case 'g':
			return KeyLowerG
		case 'h':
			return KeyLowerH
		case 'i':
			return KeyLowerI
		case 'j':
			return KeyLowerJ
		case 'k':
			return KeyLowerK
		case 'l':
			return KeyLowerL
		case 'm':
			return KeyLowerM
		case 'n':
			return KeyLowerN
		case 'o':
			return KeyLowerO
		case 'p':
			return KeyLowerP
		case 'q':
			return KeyLowerQ
		case 'r':
			return KeyLowerR
		case 's':
			return KeyLowerS
		case 't':
			return KeyLowerT
		case 'u':
			return KeyLowerU
		case 'v':
			return KeyLowerV
		case 'w':
			return KeyLowerW
		case 'x':
			return KeyLowerX
		case 'y':
			return KeyLowerY
		case 'z':
			return KeyLowerZ
		case '`':
			return KeyBacktick
		case '~':
			return KeyTilde
		case '1':
			return Key1
		case '!':
			return KeyExcelmation
		case '2':
			return Key2
		case '@':
			return KeyAt
		case '3':
			return Key3
		case '#':
			return KeyPound
		case '4':
			return Key4
		case '$':
			return KeyDollar
		case '5':
			return Key5
		case '%':
			return KeyPercent
		case '6':
			return Key6
		case '^':
			return KeyCaret
		case '7':
			return Key7
		case '&':
			return KeyAmpersand
		case '8':
			return Key8
		case '*':
			return KeyStar
		case '9':
			return Key9
		case '(':
			return KeyLeftParen
		case '0':
			return Key0
		case ')':
			return KeyRightParen
		case '-':
			return KeyMinus
		case '_':
			return KeyUnderscore
		case '=':
			return KeyEquals
		case '+':
			return KeyPlus
		case '[':
			return KeyLeftSquareBracket
		case '{':
			return KeyLeftCurlyBracket
		case ']':
			return KeyRightSquareBracket
		case '}':
			return KeyRightCurlyBracket
		case '\\':
			return KeyBackslash
		case '|':
			return KeyPipe
		case ';':
			return KeySemicolon
		case ':':
			return KeyColon
		case '\'':
			return KeySingleQuote
		case '"':
			return KeyDoubleQuote
		case ',':
			return KeyComma
		case '<':
			return KeyLessThan
		case '.':
			return KeyPeriod
		case '>':
			return KeyGreaterThan
		case '/':
			return KeyForwardSlash
		case '?':
			return KeyQuestion
		}
		switch ev.Key {
		case termbox.KeyArrowUp:
			return KeyUpArrow
		case termbox.KeyArrowLeft:
			return KeyLeftArrow
		case termbox.KeyArrowDown:
			return KeyDownArrow
		case termbox.KeyArrowRight:
			return KeyRightArrow
		case termbox.KeyEsc:
			return KeyEscape
		case termbox.KeyTab:
			return KeyTab
		case termbox.KeySpace:
			return KeySpace
		case termbox.KeyBackspace:
		case termbox.KeyBackspace2:
			return KeyBackspace
		case termbox.KeyEnter:
			return KeyReturn
		}
	}
	return KeyUnknown
}
