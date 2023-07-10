package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                    string
	AccessSecret            string
	RefreshSecret           string
	ResetSecret             string
	AccessLifetime          uint
	RefreshLifetime         uint
	ResetLifetime           uint
	EmailingAccountEmail    string
	EmailingAccountPassword string

	DatabaseUser     string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     uint
	DatabaseName     string
	DatabaseSSLMode  string
}

func STOI(envVar string) int {
	val, err := strconv.Atoi(os.Getenv(envVar))
	if err != nil {
		panic("Error converting ENV var to int")
	}
	return val
}

func NewConfig(envFile string) *Config {
	err := godotenv.Load("configs/" + envFile)
	if err != nil {
		log.Fatal("Error loading env file")
	}

	AccessLifetime := STOI("ACCESS_LIFETIME")
	RefreshLifetime := STOI("REFRESH_LIFETIME")
	ResetLifetime := STOI("RESET_LIFETIME")

	DatabasePort := STOI("DATABASE_PORT")

	return &Config{
		Port:                    os.Getenv("PORT"),
		AccessSecret:            os.Getenv("ACCESS_SECRET"),
		RefreshSecret:           os.Getenv("REFRESH_SECRET"),
		ResetSecret:             os.Getenv("RESET_SECRET"),
		AccessLifetime:          uint(AccessLifetime),
		RefreshLifetime:         uint(RefreshLifetime),
		ResetLifetime:           uint(ResetLifetime),
		EmailingAccountEmail:    os.Getenv("EMAILING_ACCOUNT_EMAIL"),
		EmailingAccountPassword: os.Getenv("EMAILING_ACCOUNT_PASSWORD"),

		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabasePort:     uint(DatabasePort),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		DatabaseSSLMode:  os.Getenv("DATABASE_SSL_MODE"),
	}
}
