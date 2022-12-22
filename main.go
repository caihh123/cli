package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/caihh123/cli/cmd"
)

var (
	app = *cli.NewApp()
)

func init() {
	app.EnableBashCompletion = true
	app.Commands = cli.Commands{
		cmd.Help,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
