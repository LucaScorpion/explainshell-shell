package cmd

import "fmt"

const ESC = "\x1B"

func Bold(text string) string {
	return modeAndReset(text, 1, 22)
}

func modeAndReset(text string, mode, reset int) string {
	return fmt.Sprintf("%s[%dm%s%s[%dm", ESC, mode, text, ESC, reset)
}
