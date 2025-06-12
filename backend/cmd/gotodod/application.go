package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const (
	DefaultConfigPath string = "/etc/gotodod"
)

// Application represents the whole App being run by main
type Application struct {
	Debug      bool
	ConfigPath string
	Config     Config
}

func NewApplication() *Application {
	return &Application{
		ConfigPath: DefaultConfigPath,
	}
}

func (*Application) Run(ctx context.Context) error {
	return nil
}

func (*Application) Close(ctx context.Context) error {
	return nil
}

// Parse command line flags and then setup the config from the
// config file
func (a *Application) LoadConfig(args []string) error {
	// Something other than the default config file?
	fs := flag.NewFlagSet("gotodod", flag.ContinueOnError)
	fs.StringVar(&a.ConfigPath, "config", DefaultConfigPath, "HTTP application config path")

	// Debugging?
	fs.BoolVar(&a.Debug, "debug", false, "show debug messages")

	if err := fs.Parse(args); err != nil {
		return err
	}

	// config path is of form '~/'?
	if strings.HasPrefix(a.ConfigPath, "~/") {
		dirname, _ := os.UserHomeDir()
		a.ConfigPath = filepath.Join(dirname, a.ConfigPath[2:])
	}

	a.DebugMessage("config directory is %s", a.ConfigPath)

	// We use Viper because life is too short
	var env string
	var ok bool
	if env, ok = os.LookupEnv("APP_ENV"); ok != true {
		return fmt.Errorf("no APP_ENV var available")
	}

	v := viper.NewWithOptions(viper.KeyDelimiter("::"))

	v.SetConfigName(env)
	v.SetConfigType("yaml")
	v.AddConfigPath(a.ConfigPath)

	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		return fmt.Errorf("fatal error config file: %s", err)
	}

	config := NewConfig()
	err = v.Unmarshal(&config)
	if err != nil {
		return fmt.Errorf("cannot unmarshall config: %s", err)
	}

	a.Config = config

	a.DebugMessage("database DSN is: %s", a.Config.DbConfig.DSN)

	return nil
}

// Print a debug message
func (a *Application) DebugMessage(format string, args ...any) {
	debugFormat := fmt.Sprintf("Debug: %s\n", format)

	if a.Debug == true {
		fmt.Fprintf(os.Stdout, debugFormat, args...)
	}
}
