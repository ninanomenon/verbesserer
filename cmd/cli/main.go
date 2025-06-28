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
			ruff := ruff.Ruff{}

			var toolSlice []tools.Tool
			toolSlice = append(toolSlice, ruff)

			tools.Run(toolSlice)

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
