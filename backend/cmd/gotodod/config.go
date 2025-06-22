package main

import (
	"net/http"
	"time"
)

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

type SessionConfig struct {
	SameSite    http.SameSite
	Secure      bool          `mapstructure:"secure"`
	CookieName  string        `mapstructure:"cookieName"`
	IdleTimeout time.Duration `mapstructure:"idleTimeout"`
	Lifetime    time.Duration `mapstructure:"lifetime"`
}

// Configuration for a HTTP server instance.
type Config struct {
	Db      DatabaseConfig `mapstructure:"database"`
	Http    HttpConfig     `mapstructure:"http"`
	Session SessionConfig  `mapstructure:"session"`
}

func NewConfig() Config {
	cfg := Config{}

	// set some defaults
	cfg.Session.SameSite = http.SameSiteStrictMode
	cfg.Session.Secure = false
	cfg.Session.CookieName = "SESSIONID"
	cfg.Session.IdleTimeout = 20 * time.Minute
	cfg.Session.Lifetime = 1 * time.Hour

	return cfg
}
