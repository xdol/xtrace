package main

import (
	"git.xdol.org/xdol/xtrace/config"
	"git.xdol.org/xdol/xtrace/router"
)

func main() {
	currentConfig := config.GetEnv()
	router.HTTPHandle(currentConfig.ListenAddress)
}
