package main

import "explainshell-shell/internal/esweb"

func main() {
	esweb.DoThing("find . -type f -print0")
}
