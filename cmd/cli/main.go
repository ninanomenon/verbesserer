package main

import (
	"context"
	"log"
	"os"

	"github.com/ninanomenon/verbesserer/internal/tools"
	"github.com/ninanomenon/verbesserer/internal/tools/ruff"
	"github.com/urfave/cli/v3"
)

func run() {
	cmd := &cli.Command{
		Name:  "Verbesserer",
		Usage: "makes it easier to make incremental improvements to your codebase!",
		Action: func(ctx context.Context, command *cli.Command) error {
			r := ruff.Ruff{}
			var t []tools.Tool
			t = append(t, r)

			tools.Run(t)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	run()
}
