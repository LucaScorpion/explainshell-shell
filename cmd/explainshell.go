package main

import (
	"explainshell-shell/internal/cmd"
	"explainshell-shell/internal/esweb"
	"fmt"
	"os"
	"strings"
)

const SEPARATOR = "──────────"

const HELP_TEXT = `
Usage: {{filename}} [options...] <command>

Options:
  --            Stop parsing options and interpret the rest as command input.
  -h, --help    Print this help text.

The CLI will stop trying to parse options as soon as it encounters a
non-option argument, or --.

Example usage:
  {{filename}} tar xzvf archive.tar.gz
`

func main() {
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

	printCommandHelp(help)
}

func printHelp() {
	fileName := os.Args[0]
	fileName = fileName[strings.LastIndexByte(fileName, '/')+1:]

	help := strings.TrimSpace(HELP_TEXT)
	help = strings.ReplaceAll(help, "{{filename}}", fileName)

	fmt.Println(help)
}

func printCommandHelp(help *esweb.CommandHelp) {
	fmt.Println("Source: " + help.Source + "\n")
	fmt.Println(help.Command + "\n")

	fmt.Println(SEPARATOR)
	for _, part := range help.Parts {
		fmt.Println(cmd.Bold(part.Part) + " " + part.ManPage + "\n")
		fmt.Println(part.Help)
		fmt.Println(SEPARATOR)
	}
}
