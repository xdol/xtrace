package main

import (
	"git.xdol.org/xdol/xtrace/internal/config"
	"git.xdol.org/xdol/xtrace/internal/router"
)

func main() {
	currentConfig := config.GetEnv()
	router.HTTPHandle(currentConfig.ListenAddress)
}
