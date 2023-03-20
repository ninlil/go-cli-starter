package main

import (
	"fmt"
	"os"

	"github.com/ninlil/go-cli-starter/src/app"
)

func main() {
	args, cmd, err := app.Init()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = cmd.Exec(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
