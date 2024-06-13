package main

import (
	"log"

	"git.xdol.org/xdol/xtrace/internal/config"
)

func main() {
	c, err := config.LoadToml("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	c.HTTP.Server()
}
