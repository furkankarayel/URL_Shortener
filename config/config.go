package config

import (
	"log"
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Configuration struct {
	Host         string `env:"HOST,required"`
	Port         string `env:"PORT,required"`
	Database     string `env:"DATABASE,required"`
	DBUser       string `env:"DB_USER,required"`
	DBPassword   string `env:"DB_PASSWORD,required"`
	IsProduction bool   `env:"IS_PRODUCTION,required"`
}

func NewConfig(files ...string) (*Configuration, error) {
	currentWorkDirectory, _ := os.Getwd()
	log.Println(currentWorkDirectory + "\\\\" + files[0])
	err := godotenv.Load(currentWorkDirectory + "\\\\" + files[0])

	if err != nil {
		log.Printf("No .env file could be found %q\n", files)
	}

	cfg := Configuration{}
	err = env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
