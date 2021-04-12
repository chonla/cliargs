package cliargs

type Cmd struct {
	Command    string
	Subcommand string
	Options    CmdOptions
	Arguments  []string
}

type CmdOptions struct {
	options map[string][]string
}
