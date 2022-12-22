package cmd

import (
	"errors"
	"fmt"

	"github.com/urfave/cli"
)

var Help = cli.Command{
	Name:   "help",
	Usage:  "<command> [arguments]",
	Action: help,
}

func help(ctx *cli.Context) error {
	switch ctx.Args().First() {
	default:
		return errors.New("unknown help topic")
	case "test":
		_, err := fmt.Fprint(ctx.App.Writer, ctx.Args())
		return err
	}
}
