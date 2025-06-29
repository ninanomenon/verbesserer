package ruff

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/ninanomenon/verbesserer/internal/tools"
)

type Ruff struct {
	CheckPath string
}

func (r Ruff) Name() string {
	return "Ruff"
}

func (r Ruff) Description() string {
	return "An extremely fast Python linter and code formatter."
}

func (r Ruff) Preflight() error {
	// TODO: Implement checks like is ruff is installed and stuff
	panic("Not implemented yet!")
}

// Run - runs `ruff check` and returns a code quality report
func (r Ruff) Run() (*[]tools.Result, error) {
	// we use the output format gitlab here to parse the json later
	ruff := exec.Command("ruff", "check", "--output-format", "gitlab", r.CheckPath)

	output, err := ruff.Output()
	// Ruff is exiting with an exit code of 1 if there are finding in the check code
	if err != nil && ruff.ProcessState.ExitCode() != 1 {
		return nil, fmt.Errorf("Ruff: unexpected error while running command: %w", err)
	}

	var report []tools.Result
	err = json.Unmarshal(output, &report)
	if err != nil {
		return nil, err
	}

	return &report, nil
}
