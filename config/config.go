package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ListenAddress string
}

func GetEnv() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	return Config{ListenAddress: os.Getenv("XTRACE_LISTEN")}

}
