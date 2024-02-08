package interaction

import "fmt"

const consoleEscape = "\x1b"

type ConsoleColor int

const (
	ConsoleColorNone ConsoleColor = iota
	ConsoleColorRed
	ConsoleColorGreen
	ConsoleColorYellow
	ConsoleColorBlue
	ConsoleColorPurple
)

func (e *Env) PrintInfoSubject(subject string, msg string) {
	fmt.Printf("%s: ", subject)
	e.printColored(ConsoleColorGreen, msg)
}

func (e *Env) printColored(color ConsoleColor, msg string) {
	fmt.Println(e.color(color) + msg + e.color(color))
}

func (e *Env) color(color ConsoleColor) string {
	if color == ConsoleColorNone {
		return fmt.Sprintf("%s[%dm", consoleEscape, color)
	}

	return fmt.Sprintf("%s[3%dm", consoleEscape, color)
}
