package configs

import "os"

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	return &Config{
		Db: DbConfig{
			DatabasePort:     os.Getenv("DATABASE_PORT"),
			DatabaseUser:     os.Getenv("DATABASE_USER"),
			DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
			DatabaseName:     os.Getenv("DATABASE_NAME"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("SECRET"),
		},
	}
}
