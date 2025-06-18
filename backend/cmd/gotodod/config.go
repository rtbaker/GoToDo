package main

// Enum type for available DB drivers. Only MySQL for now
type DatabaseDriver string

const (
	Mysql DatabaseDriver = "mysql"
)

type DatabaseConfig struct {
	Driver DatabaseDriver `mapstructure:"driver"`
	DSN    string         `mapstructure:"dsn"`
}

type HttpConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// Configuration for a HTTP server instance.
type Config struct {
	Db   DatabaseConfig `mapstructure:"database"`
	Http HttpConfig     `mapstructure:"http"`
}

func NewConfig() Config {
	return Config{}
}
