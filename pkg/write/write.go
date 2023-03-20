package pkg_write

import (
	"fmt"

	"github.com/ninlil/go-cli-starter/src/args"
)

type WriteArgs struct {
	File string `arg:"-f,--file"`
}

func (read *WriteArgs) Verify() error {
	return nil
}

func (read *WriteArgs) Exec(ma args.CommonArgs) error {
	fmt.Println("exec: Write")
	return nil
}
