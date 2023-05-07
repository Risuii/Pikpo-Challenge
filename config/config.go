package config

import (
	"fmt"
	"os"
)

type Config struct {
	App struct {
		Port string
	}
	Database struct {
		DSN string
	}
}

func New() *Config {
	c := new(Config)
	c.loadApp()
	c.loadDatabase()

	return c
}

func (c *Config) loadApp() *Config {
	port := os.Getenv("PORT")

	c.App.Port = port

	return c
}

func (c *Config) loadDatabase() *Config {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, username, password, database)

	c.Database.DSN = dsn

	return c
}
