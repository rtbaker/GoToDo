package main

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/go-viper/mapstructure/v2"
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
	SameSite    http.SameSite `mapstructure:"samesite"`
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

func HttpSamesiteFromStringViperHook() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		// Check that the data is string
		if f.Kind() != reflect.String {
			return data, nil
		}

		var samesite http.SameSite = http.SameSiteStrictMode // default strict whatever

		// Check that the target type is our custom type
		if t != reflect.TypeOf(samesite) {
			return data, nil
		}

		samesiteString := data.(string)

		switch samesiteString {
		case "SameSiteDefaultMode":
			samesite = http.SameSiteDefaultMode
		case "SameSiteLaxMode":
			samesite = http.SameSiteLaxMode
		case "SameSiteStrictMode":
			samesite = http.SameSiteStrictMode
		case "SameSiteNoneMode":
			samesite = http.SameSiteNoneMode
		}

		return samesite, nil
	}
}

func TimeDurationFromStringViperHook() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{},
	) (interface{}, error) {
		// Check that the data is string
		if f.Kind() != reflect.String {
			return data, nil
		}

		var duration time.Duration = 20 * time.Minute // reasonable default?

		// Check that the target type is our custom type
		if t != reflect.TypeOf(duration) {
			return data, nil
		}

		pattern, err := regexp.Compile(`(\d+)([smhd])`)

		if err != nil {
			return nil, fmt.Errorf("TimeDurationFromStringViperHook: %s", err)
		}

		durationString := data.(string)

		matches := pattern.FindStringSubmatch(durationString)

		if matches == nil || len(matches) != 3 {
			return nil, errors.New("durations should be specified as length[smhd], e.g. 10s")
		}

		durationValue, _ := strconv.Atoi(matches[1]) // ignore error as the regex ensures it's an int

		switch matches[2] {
		case "s":
			duration = time.Duration(durationValue) * time.Second
		case "m":
			duration = time.Duration(durationValue) * time.Minute
		case "h":
			duration = time.Duration(durationValue) * time.Hour
		case "d":
			duration = time.Duration(durationValue) * time.Hour * 24
		}

		return duration, nil
	}
}
