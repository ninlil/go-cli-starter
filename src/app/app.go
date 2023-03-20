package app

import (
	"fmt"

	"github.com/alexflint/go-arg"
)

func Init() (*cmdArgs, subcommand, error) {
	pa := arg.MustParse(&myArgs)
	if pa == nil {
		return nil, nil, fmt.Errorf("arg weird failed")
	}

	var subcmd subcommand

	switch true {
	case myArgs.Read != nil:
		subcmd = myArgs.Read
	case myArgs.Write != nil:
		subcmd = myArgs.Write
	default:
		return nil, nil, fmt.Errorf("command missing")
	}

	return nil, subcmd, subcmd.Verify()
}
