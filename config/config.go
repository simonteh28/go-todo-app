package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	port string
	dbUser	string
	dbPass	string
	dbHost	string
	dbName 	string
	dbPort	string
}

func Get() (*Config, error) {
	conf := &Config{}

	// APP_PORT
	flag.StringVar(&conf.port, "port", os.Getenv("PORT"), "Port")
	if conf.port == "" {
		return nil, errors.New("invalid configuration. environment variable PORT not found")
	}

	// POSTGRES_USER
	flag.StringVar(&conf.dbUser, "dbuser", os.Getenv("POSTGRES_USER"), "DB User name")
	if conf.dbUser == "" {
		return nil, errors.New("invalid configuration. environment variable POSTGRES_USER not found")
	}
	
	// POSTGRES_PASSWORD
	flag.StringVar(&conf.dbPass, "dbpass", os.Getenv("POSTGRES_PASSWORD"), "DB Password")
	if conf.dbPass == "" {
		return nil, errors.New("invalid configuration. environment variable POSTGRES_PASSWORD not found")
	}

	// POSTGRES_HOST
	flag.StringVar(&conf.dbHost, "dbhost", os.Getenv("POSTGRES_HOST"), "DB Host")
	if conf.dbHost == "" {
		return nil, errors.New("invalid configuration. environment variable POSTGRES_HOST not found")
	}

	// POSTGRES_DB
	flag.StringVar(&conf.dbName, "dbname", os.Getenv("POSTGRES_DB"), "DB Name")
	if conf.dbName == "" {
		return nil, errors.New("invalid configuration. environment variable POSTGRES_DB not found")
	}
	
	// POSTGRES_PORT
	flag.StringVar(&conf.dbPort, "dbport", os.Getenv("POSTGRES_PORT"), "DB Port")
	if conf.dbPort == "" {
		return nil, errors.New("invalid configuration. environment variable POSTGRES_PORT not found")
	}

	flag.Parse()

	return conf, nil
}

func (c *Config) GetDBConnString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.dbUser,
		c.dbPass,
		c.dbHost,
		c.dbPort,
		c.dbName,
	)

	// return "postgres://postgres:postgres@localhost:5432/todo-app?sslmode=disable"
}

func (c *Config) GetDBString() string {
	return fmt.Sprintf("%s:%s/%s", c.dbHost, c.dbPort, c.dbName)
}
