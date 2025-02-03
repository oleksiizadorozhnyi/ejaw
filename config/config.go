package config

import "github.com/caarlos0/env/v9"

var (
	config = Config{}
)

type Config struct {
	ServerPort string `env:"SERVER_PORT" envDefault:":8080"`
	Postgres   Postgres
	Admin      Admin
}

type Postgres struct {
	User     string `env:"DB_USER" envDefault:"postgres"`
	Password string `env:"DB_PASSWORD" envDefault:"password"`
	Name     string `env:"DB_NAME" envDefault:"shopdb"`
	Host     string `env:"DB_HOST" envDefault:"localhost"`
}

type Admin struct {
	User     string `env:"ADMIN_USERNAME" envDefault:"admin"`
	Password string `env:"ADMIN_PASSWORD" envDefault:"123"`
}

func LoadConfig() (Config, error) {
	err := env.Parse(&config)
	return config, err
}
