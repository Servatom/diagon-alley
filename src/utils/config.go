package utils

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort         string `env:"SERVER_PORT" envDefault:"3000"`
	DatabaseName       string `env:"DATABASE_NAME" envDefault:"polyjuice"`
	DatabaseHost       string `env:"DATABASE_HOST" envDefault:"localhost"`
	DatabasePort       int    `env:"DATABASE_PORT" envDefault:"5432"`
	DatabaseUser       string `env:"DATABASE_USER" envDefault:"postgres"`
	DatabasePassword   string `env:"DATABASE_PASSWORD" envDefault:"postgres"`
	DatabaseSSLMode    string `env:"DATABASE_SSL_MODE" envDefault:"disable"`
	Timezone           string `env:"TIMEZONE" envDefault:"Asia/Kolkata"`
	SecretKey          string `env:"SECRET_KEY" envDefault:"secret"`
	GoogleClientID     string `env:"GOOGLE_CLIENT_ID" envDefault:""`
	GoogleClientSecret string `env:"GOOGLE_CLIENT_SECRET" envDefault:""`
	GoogleRedirectURL  string `env:"GOOGLE_REDIRECT_URL" envDefault:"http://localhost:9000"`
}

func NewConfig() *Config {
	error := godotenv.Load()
	if error != nil {
		log.Print("Error loading .env file")
	}
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg
}

func GetPostgresConnectionString(cfg *Config) string {
	psqlURL := fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v",
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName,
	)
	return psqlURL
}
