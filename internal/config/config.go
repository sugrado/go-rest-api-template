package config

import "os"

var (
	configuration *Configuration
)

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type ServerConfiguration struct {
	Environment            string
	Port                   string
	Domain                 string
	SecretKey              string
	AccessTokenExpiration  string
	RefreshTokenExpiration string
}

type DatabaseConfiguration struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

func Load() (*Configuration, error) {
	configuration = &Configuration{
		Server: ServerConfiguration{
			Port:                   os.Getenv("SERVER_PORT"),
			Domain:                 os.Getenv("SERVER_DOMAIN"),
			Environment:            os.Getenv("SERVER_ENV"),
			SecretKey:              os.Getenv("SERVER_SECRET"),
			AccessTokenExpiration:  os.Getenv("SERVER_ACCESS_TOKEN_EXPIRATION"),
			RefreshTokenExpiration: os.Getenv("SERVER_REFRESH_TOKEN_EXPIRATION"),
		},
		Database: DatabaseConfiguration{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}
	return configuration, nil
}
