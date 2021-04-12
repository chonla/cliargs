package cliargs

import (
	"errors"
	"os"
	"strings"
)

func Parse() (Cmd, error) {
	argCommand := ""
	argSubcommand := ""
	argOptions := CmdOptions{
		options: map[string][]string{},
	}
	argArguments := []string{}

	argLen := len(os.Args)

	// Command and Subcommand are positional
	if argLen > 1 {
		if argLen > 2 {
			argSubcommand = os.Args[2]
		}
		argCommand = os.Args[1]
	}

	for argIndex := 1; argIndex < argLen; argIndex++ {
		switch argIndex {
		case 1:
			if isOptional(os.Args[argIndex]) {
				return Cmd{}, errors.New("Command is required")
			}
			argCommand = os.Args[argIndex]
		case 2:
			if isOptional(os.Args[argIndex]) {
				argOptions.Add(os.Args[argIndex])
			} else {
				argSubcommand = os.Args[argIndex]
			}
		default:
			if isOptional(os.Args[argIndex]) {
				argOptions.Add(os.Args[argIndex])
			} else {
				argArguments = append(argArguments, os.Args[argIndex])
			}
		}
	}

	return Cmd{
		Command:    argCommand,
		Subcommand: argSubcommand,
		Options:    argOptions,
		Arguments:  argArguments,
	}, nil
}

func isOptional(v string) bool {
	return strings.HasPrefix(v, "-")
}
