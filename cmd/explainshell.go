package main

import (
	"explainshell-shell/internal/cmd"
	"explainshell-shell/internal/esweb"
	"fmt"
	"os"
	"strings"
)

const separator = "──────────"

const helpText = `
Usage: {{filename}} [options...] <command>

Options:
  --                Stop parsing options and interpret the rest as command input.
  -h, --help        Print this help text.
  -n, --no-color    Disable color output.

The CLI will stop trying to parse options as soon as it encounters a
non-option argument, or --.

Example usage:
  {{filename}} tar xzvf archive.tar.gz
`

var COLORS = []int{32, 33, 34, 35, 36}

func main() {
	colors := true
	argsI := 1
	for ; argsI < len(os.Args); argsI++ {
		arg := os.Args[argsI]

		if len(arg) == 0 || arg == "--" {
			argsI++
			break
		}

		if arg[0] != '-' {
			break
		}

		switch os.Args[argsI] {
		case "-h":
			fallthrough
		case "--help":
			printHelp()
			return
		case "-n":
			fallthrough
		case "--no-color":
			colors = false
		default:
			fmt.Println("Invalid option: " + arg)
			printHelp()
			os.Exit(1)
		}
	}

	input := ""
	for ; argsI < len(os.Args); argsI++ {
		input += os.Args[argsI] + " "
	}
	input = strings.TrimSpace(input)

	if len(input) == 0 {
		fmt.Println("No command given")
		printHelp()
		os.Exit(1)
	}

	help, err := esweb.GetCommandHelp(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	printCommandHelp(help, colors)
}

func printHelp() {
	fileName := os.Args[0]
	fileName = fileName[strings.LastIndexByte(fileName, '/')+1:]

	help := strings.TrimSpace(helpText)
	help = strings.ReplaceAll(help, "{{filename}}", fileName)

	fmt.Println(help)
}

func printCommandHelp(help *esweb.CommandHelp, colors bool) {
	fmt.Println("Source: " + help.Source + "\n")

	if len(help.Parts) == 0 {
		fmt.Println(cmd.Bold("No man page found."))
		return
	}

	colorIndex := 0
	for i, part := range help.Parts {
		if i == 0 {
			fmt.Print(cmd.Bold(part.Part))
		} else {
			fmt.Print(cmd.Bold(colorText(part.Part, COLORS[colorIndex], colors)))
			colorIndex = (colorIndex + 1) % len(COLORS)
		}

		if i < len(help.Parts)-1 {
			fmt.Print(" ")
		}
	}

	fmt.Println("\n\n" + separator)
	colorIndex = 0
	for i, part := range help.Parts {
		if i == 0 {
			fmt.Println(cmd.Bold(part.Part) + " " + help.ManPage + "\n")
		} else {
			fmt.Println(cmd.Bold(colorText(part.Part, COLORS[colorIndex], colors)) + "\n")
			colorIndex = (colorIndex + 1) % len(COLORS)
		}

		fmt.Println(part.Help)
		fmt.Println(separator)
	}
}

func colorText(text string, color int, colors bool) string {
	if !colors {
		return text
	}
	return cmd.Color(text, color)
}
