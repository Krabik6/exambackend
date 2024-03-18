package config

import (
	"github.com/joho/godotenv"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	User     string `envconfig:"POSTGRES_USER" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Host     string `envconfig:"POSTGRES_HOST" required:"true"`
	Port     string `envconfig:"POSTGRES_PORT" required:"true"`
	DBName   string `envconfig:"POSTGRES_DB" required:"true"`
}

type Config struct {
	DBConfig DBConfig
}

var (
	once sync.Once
	cfg  Config
)

func New() *Config {
	once.Do(func() {
		path := ".env"
		godotenv.Load(path)
		if err := envconfig.Process("", &cfg); err != nil {
			panic(err)
		}
	})
	return &cfg
}
