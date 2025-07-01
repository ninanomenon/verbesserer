package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

const configFileName = ".verbesserer.toml"

const defaultCheckPath = "./"

type Config struct {
	Ruff Ruff `toml:"ruff,omitempty"`
}

func (c *Config) fillEmptyWithDefaults() {
	c.Ruff.fillEmptyWithDefaults()
}

type Ruff struct {
	// The path of the code that should be checked
	// defaults to the current working directory.
	CheckPath string `toml:"check_path,omitempty"`
}

func (r *Ruff) fillEmptyWithDefaults() {
	if r.CheckPath == "" {
		r.CheckPath = defaultCheckPath
	}
}

func LoadConfig(path string) (Config, error) {
	filePath := configFileName
	if path != "" {
		filePath = filepath.Join(path, configFileName)
	}

	var config Config

	file, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			config.fillEmptyWithDefaults()

			return config, nil
		}

		return Config{}, fmt.Errorf("load config: %w", err)
	}

	err = toml.Unmarshal(file, &config)
	if err != nil {
		return Config{}, fmt.Errorf("load config: %w", err)
	}

	config.fillEmptyWithDefaults()

	return config, nil
}
