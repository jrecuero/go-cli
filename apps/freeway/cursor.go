package freeway

import "fmt"

// Esc represents ...
var Esc = "\x1b"

func escape(format string, args ...interface{}) string {
	return fmt.Sprintf("%s%s", Esc, fmt.Sprintf(format, args...))
}

// Show returns ANSI escape sequence to show the cursor
func Show() string {
	return escape("[?25h")
}

// Hide returns ANSI escape sequence to hide the cursor
func Hide() string {
	return escape("[?25l")
}

// MoveTo returns ANSI escape sequence to move cursor
// to specified position on screen.
func MoveTo(line, col int) string {
	return escape("[%d;%dH", line, col)
}

// MoveUp returns ANSI escape sequence to move cursor
// up n lines.
func MoveUp(n int) string {
	return escape("[%dA", n)
}

// MoveDown returns ANSI escape sequence to move cursor
// down n lines.
func MoveDown(n int) string {
	return escape("[%dB", n)
}

// MoveRight returns ANSI escape sequence to move cursor
// right n columns.
func MoveRight(n int) string {
	return escape("[%dC", n)
}

// MoveLeft returns ANSI escape sequence to move cursor
// left n columns.
func MoveLeft(n int) string {
	return escape("[%dD", n)
}

// MoveUpperLeft returns ANSI escape sequence to move cursor
// to upper left corner of screen.
func MoveUpperLeft(n int) string {
	return escape("[%dH", n)
}

// MoveNextLine returns ANSI escape sequence to move cursor
// to next line.
func MoveNextLine() string {
	return escape("E")
}

// ClearLineRight returns ANSI escape sequence to clear line
// from right of the cursor.
func ClearLineRight() string {
	return escape("[0K")
}

// ClearLineLeft returns ANSI escape sequence to clear line
// from left of the cursor.
func ClearLineLeft() string {
	return escape("[1K")
}

// ClearEntireLine returns ANSI escape sequence to clear current line.
func ClearEntireLine() string {
	return escape("[2K")
}

// ClearScreenDown returns ANSI escape sequence to clear screen below.
// the cursor.
func ClearScreenDown() string {
	return escape("[0J")
}

// ClearScreenUp returns ANSI escape sequence to clear screen above.
// the cursor.
func ClearScreenUp() string {
	return escape("[1J")
}

// ClearEntireScreen returns ANSI escape sequence to clear the screen.
func ClearEntireScreen() string {
	return escape("[2J")
}

// SaveAttributes returns ANSI escape sequence to save current position
// and attributes of the cursor.
func SaveAttributes() string {
	return escape("7")
}

// RestoreAttributes returns ANSI escape sequence to restore saved position
// and attributes of the cursor.
func RestoreAttributes() string {
	return escape("8")
}

// Reset is ...
func Reset() string {
	return escape("[0m")
}

// BoldOn is ...
func BoldOn() string {
	return escape("[1m")
}

// ItalicsOn is ...
func ItalicsOn() string {
	return escape("[3m")
}

// UnderlineOn is ...
func UnderlineOn() string {
	return escape("[4m")
}

// InverseOn is ...
func InverseOn() string {
	return escape("[7m")
}

// StrikeThroughOn is ...
func StrikeThroughOn() string {
	return escape("[9m")
}

// BoldOff is ...
func BoldOff() string {
	return escape("[22m")
}

// ItalicsOff is ...
func ItalicsOff() string {
	return escape("[23m")
}

// UnderlineOff is ...
func UnderlineOff() string {
	return escape("[24m")
}

// InverseOff is ...
func InverseOff() string {
	return escape("[27m")
}

// StrikeThroughOff is ...
func StrikeThroughOff() string {
	return escape("[29m")
}

// FgRed is ...
func FgRed() string {
	return escape("[31m")
}

// FgGree is ...
func FgGree() string {
	return escape("[32m")
}

// FgYellow is ...
func FgYellow() string {
	return escape("[33m")
}

// FgBlue is ...
func FgBlue() string {
	return escape("[34m")
}

// FgMagenta is ...
func FgMagenta() string {
	return escape("[35m")
}

// FgCyan is ...
func FgCyan() string {
	return escape("[36m")
}

// FgWhite is ...
func FgWhite() string {
	return escape("[37m")
}

// FgDefault is ...
func FgDefault() string {
	return escape("[39m")
}

// FgBlack is ...
func FgBlack() string {
	return escape("[40m")
}

// BgRed is ...
func BgRed() string {
	return escape("[41m")
}

// BgGreen is ...
func BgGreen() string {
	return escape("[42m")
}

// BgYellow is ...
func BgYellow() string {
	return escape("[43m")
}

// BgBlue is ...
func BgBlue() string {
	return escape("[44m")
}

// BgMagenta is ...
func BgMagenta() string {
	return escape("[45m")
}

// BgCyan is ...
func BgCyan() string {
	return escape("[46m")
}

// BgWhite is ...
func BgWhite() string {
	return escape("[47m")
}

// BgDefault is ...
func BgDefault() string {
	return escape("[49m")
}
