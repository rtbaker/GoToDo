package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/rtbaker/GoToDo/database/mysql"
	"github.com/rtbaker/GoToDo/http"
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
	HTTPServer *http.Server
	DB         *sql.DB
}

func NewApplication() *Application {
	return &Application{
		ConfigPath: DefaultConfigPath,
	}
}

func (a *Application) Run(ctx context.Context) error {
	// Connect to the DB
	a.getDBConnection()

	// Setup an HTTP Server
	a.HTTPServer = http.NewServer()
	a.HTTPServer.Host = a.Config.Http.Host
	a.HTTPServer.Port = a.Config.Http.Port

	err := a.HTTPServer.Run()
	if err != nil {
		return fmt.Errorf("error running http server: %s", err)
	}

	return nil
}

func (a *Application) Close() error {
	if a.HTTPServer != nil {
		if err := a.HTTPServer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Work out which DB driver we are using and get the right db connection
// (Only mysql for now)
func (a *Application) getDBConnection() error {
	switch a.Config.Db.Driver {
	case Mysql:
		var err error
		a.DB, err = mysql.NewDB(a.Config.Db.DSN)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown DB driver in config: %s", a.Config.Db.Driver)
	}

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

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(a.ConfigPath)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("fatal error config file: %s", err)
	}

	config := NewConfig()

	// Do this so we can use ENV vars in the yaml
	for _, k := range viper.AllKeys() {
		v := viper.GetString(k)
		viper.Set(k, os.ExpandEnv(v))
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return fmt.Errorf("cannot unmarshall config: %s", err)
	}

	a.Config = config

	a.DebugMessage("database DSN is: %s", a.Config.Db.DSN)
	a.DebugMessage("http listen on: \"%s:%d\"", a.Config.Http.Host, a.Config.Http.Port)

	return nil
}

// Print a debug message
func (a *Application) DebugMessage(format string, args ...any) {
	debugFormat := fmt.Sprintf("Debug: %s\n", format)

	if a.Debug == true {
		fmt.Fprintf(os.Stdout, debugFormat, args...)
	}
}
