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

func (r *Ruff) Name() string {
	return "Ruff"
}

func (r *Ruff) Description() string {
	return "An extremely fast Python linter and code formatter."
}

// Run - runs `ruff check` and returns a code quality report
func (r *Ruff) Run() (*[]tools.Result, error) {
	path, err := exec.LookPath("ruff")
	if err != nil {
		return nil, err
	}

	// we use the output format gitlab here to parse the json later
	ruff := exec.Command(path, "check", "--output-format", "gitlab", r.CheckPath)

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
