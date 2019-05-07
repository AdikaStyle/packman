package controllers

import (
	"errors"
	"fmt"
	"gopkg.in/urfave/cli.v2"
)

var UnpackCommand = cli.Command{
	Name:      "unpack",
	Aliases:   []string{"u"},
	Usage:     "unpack <package-name> <dest-path>",
	UsageText: "Will unpack the given package to the given destination path",
	Action: func(context *cli.Context) error {

		if context.NArg() != 2 {
			fmt.Println("unpack expects exactly 2 argument")
		}

		packageName := context.Args().Get(0)
		path := context.Args().Get(1)

		if packageName == "" {
			return errors.New("you must provide a package name")
		}

		if path == "" {
			return errors.New("you must provide a path to the project")
		}

		return PackmanModule.Unpacker.Unpack(packageName, path, flagsArray(context))
	},
}

func flagsArray(ctx *cli.Context) []string {
	out := make([]string, ctx.NumFlags())
	for _, flagName := range ctx.FlagNames() {
		out = append(out, flagName, ctx.String(flagName))
	}

	return out
}
