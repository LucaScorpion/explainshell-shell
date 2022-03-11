package esweb

type CommandHelp struct {
	Command string
	Parts   []*CommandPart
}

type CommandPart struct {
	Part string
	Help string
}
