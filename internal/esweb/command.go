package esweb

type CommandHelp struct {
	Parts []*CommandPart
}

type CommandPart struct {
	Part string
	Help string
}
