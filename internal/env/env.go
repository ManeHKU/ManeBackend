package env

import (
	"fmt"
	"github.com/caarlos0/env/v10"
	"log"
)

type Config struct {
	PORT          uint   `env:"PORT" envDefault:"8080"`
	MACHINE_NAME  string `env:"K_REVISION" envDefault:"LOCAL"`
	DATABASE_URI  string `env:"DB_URI"`
	DATABASE_CERT string `env:"DB_CERT"`
	JWT_SECRET    string `env:"JWT_SECRET"`
}

var config *Config

func loadEnv() {
	options := env.Options{RequiredIfNoDef: true}

	if err := env.ParseWithOptions(config, options); err != nil {
		log.Fatal(err)
	}

	log.SetPrefix(fmt.Sprintf("[%v]\t", config.MACHINE_NAME))
	log.Printf("Loaded all environment variables, PORT: %v", config.PORT)
}

func GetConfig() *Config {
	if config == nil {
		loadEnv()
	}
	return config
}
