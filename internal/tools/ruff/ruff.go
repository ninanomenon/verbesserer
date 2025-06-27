package ruff

import (
	"os/exec"
)

type Ruff struct{}

func (r Ruff) Name() string {
	return "Ruff"
}

func (r Ruff) Description() string {
	return "Ruff python bla foo"
}

func (r Ruff) Run() (string, error) {
	// we use the output format gitlab here to parse the json later
	ruff := exec.Command("ruff", "check", "--output-format", "gitlab")

	output, err := ruff.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
