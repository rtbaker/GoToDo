package main

type DatabaseConfig struct {
	DSN string `mapstructure:"dsn"`
}

// Configuration for a HTTP server instance.
type Config struct {
	DbConfig DatabaseConfig `mapstructure:"database"`
}

func NewConfig() Config {
	return Config{}
}
