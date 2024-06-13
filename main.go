package main

import (
	"log"
	"os"

	"git.xdol.org/xdol/xtrace/internal/config"
)

func main() {
	var configPath string
	switch len(os.Args) {
	case 2: //nolint:mnd
		configPath = os.Args[1]
	case 1:
		configPath = "config.toml"
	default:
		log.Fatal("Max 1 argument is valid.")
	}

	c, err := config.LoadToml(configPath)
	if err != nil {
		log.Fatal(err)
	}

	c.HTTP.Server()
}
