package cmd

import "fmt"

const esc = "\x1B"

func Bold(text string) string {
	return modeAndReset(text, 1, 22)
}

func Color(text string, color int) string {
	return modeAndReset(text, color, 39)
}

func modeAndReset(text string, mode, reset int) string {
	return fmt.Sprintf("%s[%dm%s%s[%dm", esc, mode, text, esc, reset)
}
