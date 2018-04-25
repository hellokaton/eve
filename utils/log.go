package utils

import (
	"fmt"
	"time"
)

const (
	color_red     = uint8(iota + 91)
	color_green
	color_yellow
	color_blue
	color_magenta  //洋红
	info          = "[INFO]"
	trac          = "[TRAC]"
	erro          = "[ERRO]"
	warn          = "[WARN]"
	succ          = "[SUCC]"
)

// see complete color rules in document in https://en.wikipedia.org/wiki/ANSI_escape_code#cite_note-ecma48-13
func Trace(format string, a ...interface{}) {
	prefix := Yellow(trac)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}

func Info(format string, a ...interface{}) {
	prefix := Blue(info)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}

func Success(format string, a ...interface{}) {
	prefix := Green(succ)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}

func Warning(format string, a ...interface{}) {
	prefix := Magenta(warn)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}

func Error(format string, a ...interface{}) {
	prefix := Red(erro)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}

func Red(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_red, s)
}

func Green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_green, s)
}

func Yellow(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_yellow, s)
}

func Blue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_blue, s)
}

func Magenta(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_magenta, s)
}

func formatLog(prefix string) string {
	return time.Now().Format("2006/01/02 15:04:05") + " " + prefix + " "
}
