package pkg_read

import (
	"fmt"

	args "github.com/ninlil/go-cli-starter/src/args"
)

type ReadArgs struct {
	File string `arg:"-f,--file"`
}

func (read *ReadArgs) Verify() error {
	return nil
}

func (read *ReadArgs) Exec(ma args.CommonArgs) error {
	fmt.Println("exec: Read")
	return nil
}
