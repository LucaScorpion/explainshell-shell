package esweb

type CommandHelp struct {
	Source  string
	Command string
	ManPage string
	Parts   []*CommandPart
}

type CommandPart struct {
	Part string
	Help string
}
