package main

type DatabaseConfig struct {
	DSN string `mapstructure:"dsn"`
}

type HttpConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure: "port"`
}

// Configuration for a HTTP server instance.
type Config struct {
	Db   DatabaseConfig `mapstructure:"database"`
	Http HttpConfig     `mapstructure:"http"`
}

func NewConfig() Config {
	return Config{}
}
