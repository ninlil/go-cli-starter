package app

import (
	args "github.com/ninlil/go-cli-starter/src/args"

	pkg_read "github.com/ninlil/go-cli-starter/pkg/read"
	pkg_write "github.com/ninlil/go-cli-starter/pkg/write"
)

type subcommand interface {
	Verify() error
	Exec(args.CommonArgs) error
}

type cmdArgs struct {
	VerboseFlag bool                 `arg:"-v,--verbose" help:"Verbose mode (more logging)"`
	Read        *pkg_read.ReadArgs   `arg:"subcommand:read"`
	Write       *pkg_write.WriteArgs `arg:"subcommand:write"`
}

func (arg *cmdArgs) Description() string {
	return "cmdArgs-description"
}

func (arg *cmdArgs) Version() string {
	return "v?.?"
}

func (arg *cmdArgs) Verbose() bool {
	return arg.VerboseFlag
}

func (arg *cmdArgs) Exec(subcmd subcommand) {
	subcmd.Exec(arg)
}

var (
	myArgs cmdArgs
)
