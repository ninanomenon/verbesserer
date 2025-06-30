package main

import (
	"context"
	"log"
	"os"

	"github.com/ninanomenon/verbesserer/cmd/cli/commands"
	"github.com/urfave/cli/v3"
)

var cmd *cli.Command

func init() {
	cmd = &cli.Command{
		Name:     "Verbesserer",
		Usage:    "makes it easier to make incremental improvements to your codebase!",
		Commands: []*cli.Command{commands.CheckCommand()},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config-path",
				Aliases: []string{"cp"},
				Value:   "",
				Usage:   "Customize path to the config file (.verbesserer.toml). Default: Current working directory.",
			},
		},
	}
}

func run() {
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	run()
}
