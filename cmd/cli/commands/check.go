package commands

import (
	"context"

	"github.com/ninanomenon/verbesserer/internal/tools"
	"github.com/ninanomenon/verbesserer/internal/tools/ruff"
	"github.com/urfave/cli/v3"
)

func CheckCommand() *cli.Command {
	return &cli.Command{
		Name:        "run",
		Usage:       "",
		UsageText:   "",
		Description: "",
		Action:      runAction,
	}
}

func runAction(ctx context.Context, command *cli.Command) error {
	ruff := ruff.Ruff{
		CheckPath: "internal/tools/ruff/test_data/",
	}

	var toolSlice []tools.Tool
	toolSlice = append(toolSlice, ruff)

	tools.Run(toolSlice)

	return nil
}
