package esweb

type CommandHelp struct {
	Source  string
	Command string
	Parts   []*CommandPart
}

type CommandPart struct {
	Part    string
	Help    string
	ManPage string
}
