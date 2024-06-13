package config

import (
	"os"

	"git.xdol.org/xdol/xtrace/internal/router"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	HTTP router.Config
}

func LoadToml(file string) (Config, error) {
	var c Config

	source, err := os.ReadFile(file)
	if err != nil {
		return c, errConfigFileNotReadable
	}

	err = toml.Unmarshal(source, &c)
	if err != nil {
		panic(err)
	}

	return c, nil
}
