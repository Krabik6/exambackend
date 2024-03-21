package config

import (
	"github.com/joho/godotenv"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     string `envconfig:"DB_PORT_INTERNAL" required:"true"`
	DBName   string `envconfig:"DB_NAME" required:"true"`
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
