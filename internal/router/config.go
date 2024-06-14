package router

import "git.xdol.org/xdol/xtrace/internal/router/errors"

type Config struct {
	ListenAddress string `toml:"listen"`
	Error         errors.Config
}
