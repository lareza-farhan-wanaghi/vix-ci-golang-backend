package config

import (
	"os"
	"self-payrol/config/postgres"
	"strconv"

	"gorm.io/gorm"
)

type (
	config struct {
		DB *gorm.DB
	}

	Config interface {
		ServiceName() string
		ServicePort() int
		ServiceEnvironment() string
		Database() *gorm.DB
	}
)

// NewConfig returns a new config object
func NewConfig() Config {
	return &config{
		DB: postgres.InitGorm(),
	}
}

// Database returns the database of the config object
func (c *config) Database() *gorm.DB {
	return c.DB
}

// ServiceName returns the service name of the config object
func (c *config) ServiceName() string {
	return os.Getenv("SERVICE_NAME")
}

// ServicePort returns the service port of the config object
func (c *config) ServicePort() int {
	v := os.Getenv("PORT")
	port, _ := strconv.Atoi(v)

	return port
}

// ServiceEnvironment returns the service environment data of the config object
func (c *config) ServiceEnvironment() string {
	return os.Getenv("ENV")
}
