package report

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

const userWrite = 600

type Report struct {
	FileHash string  `toml:"file_hash"`
	Issues   []Issue `toml:"issues"`
}

type Issue struct {
	Message string `toml:"message"`
	Hash    string `toml:"hash"`
	Lines   Lines  `toml:"lines"`
}

type Lines struct {
	Begin int `toml:"begin"`
	End   int `toml:"end,omitempty"`
}

type Reports map[string]Report

func (r Reports) WriteToml() error {
	bytes, err := toml.Marshal(r)
	if err != nil {
		return err
	}

	err = os.WriteFile(".verbesserer.result", bytes, userWrite)
	if err != nil {
		return err
	}

	return nil
}
