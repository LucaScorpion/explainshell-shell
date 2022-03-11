package main

import (
	"explainshell-shell/internal/cmd"
	"explainshell-shell/internal/esweb"
	"fmt"
	"os"
)

const SEPARATOR = "──────────"

func main() {
	help, err := esweb.GetCommandHelp("ssh -i keyfile -f -N -L 1234:www.google.com:80 host")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	printCommandHelp(help)
}

func printCommandHelp(help *esweb.CommandHelp) {
	fmt.Println("\n" + help.Command + "\n")

	fmt.Println(SEPARATOR)
	for _, part := range help.Parts {
		fmt.Println(cmd.Bold(part.Part) + "\n")
		fmt.Println(part.Help)
		fmt.Println(SEPARATOR)
	}
}
