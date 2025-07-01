package commands

import (
	"context"
	"fmt"

	"github.com/ninanomenon/verbesserer/internal"
	"github.com/ninanomenon/verbesserer/internal/config"
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
	configPath := command.String("config-path")
	config, err := config.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("run action: %w", err)
	}

	ruff := &ruff.Ruff{
		CheckPath: config.Ruff.CheckPath,
	}

	var toolSlice []tools.Tool
	toolSlice = append(toolSlice, ruff)

	report, errors := internal.Execute(toolSlice)
	if len(errors) != 0 {
		for _, error := range errors {
			fmt.Println(error)
		}
	}

	for _, r := range report {
		fmt.Printf("Reports: %#v\n", r)
		for _, rr := range r.Issues {
			fmt.Printf("Issue: %#v\n", rr)
		}
	}

	err = report.WriteToml()
	if err != nil {
		return fmt.Errorf("run action: %w", err)
	}

	return nil
}
