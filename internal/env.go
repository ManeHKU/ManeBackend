package internal

import (
	"fmt"
	"github.com/caarlos0/env/v10"
	"log"
)

type Config struct {
	PORT         uint   `env:"PORT" envDefault:"8080"`
	MACHINE_NAME string `env:"K_REVISION" envDefault:"LOCAL"`
	DB_HOST      string `env:"DB_HOST"`
	DB_PASSWORD  string `env:"DB_PASSWORD"`
	JWT_SECRET   string `env:"JWT_SECRET"`
}

func LoadEnv() *Config {
	config := &Config{}
	options := env.Options{RequiredIfNoDef: true}

	if err := env.ParseWithOptions(config, options); err != nil {
		log.Fatal(err)
	}

	log.SetPrefix(fmt.Sprintf("[%v]\t", config.MACHINE_NAME))
	log.Printf("Loaded all environment variables, PORT: %v", config.PORT)
	return config
}
