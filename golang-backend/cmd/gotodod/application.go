package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
	"github.com/go-viper/mapstructure/v2"
	gotodo "github.com/rtbaker/GoToDo/Model"
	"github.com/rtbaker/GoToDo/database/inmemory"
	"github.com/rtbaker/GoToDo/database/mysql"
	"github.com/rtbaker/GoToDo/http"
	"github.com/spf13/viper"
)

const (
	DefaultConfigPath string = "/etc/gotodod"
)

// Application represents the whole App being run by main
type Application struct {
	Debug        bool
	ConfigPath   string
	Config       Config
	HTTPServer   *http.Server
	DB           *sql.DB
	TodoService  gotodo.ToDoService
	UserService  gotodo.UserService
	SessionStore scs.Store
	Logger       *log.Logger
}

func NewApplication() *Application {
	return &Application{
		ConfigPath: DefaultConfigPath,
	}
}

func (a *Application) Run(ctx context.Context) error {
	// Connect to the DB
	err := a.GetDBService()
	if err != nil {
		return fmt.Errorf("error getting DB service: %s", err)
	}

	sessionCfg := http.SessionConfig{
		IdleTimeout: a.Config.Session.IdleTimeout,
		Lifetime:    a.Config.Session.Lifetime,
		Name:        a.Config.Session.CookieName,
		SameSite:    a.Config.Session.SameSite,
		Secure:      a.Config.Session.Secure,
	}

	// Setup a logger, for now we just use stdout
	a.Logger = log.New(os.Stdout, "gotodod:", log.LstdFlags)

	// Setup an HTTP Server
	a.HTTPServer = http.NewServer(sessionCfg)
	a.HTTPServer.Host = a.Config.Http.Host
	a.HTTPServer.Port = a.Config.Http.Port

	a.HTTPServer.UserService = a.UserService
	a.HTTPServer.TodoService = a.TodoService
	a.HTTPServer.SessionManager.Store = a.SessionStore

	a.HTTPServer.Logger = a.Logger

	// And run our server
	err = a.HTTPServer.Run()
	if err != nil {
		return fmt.Errorf("error running http server: %s", err)
	}

	return nil
}

func (a *Application) Close() error {
	if a.HTTPServer != nil {
		if err := a.HTTPServer.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "error closing http server: %s\n", err)
		}
	}

	if a.DB != nil {
		if err := a.DB.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "error closing DB server: %s\n", err)
		}
	}
	return nil
}

// Work out which DB driver we are using and get the right db connection
// (Only mysql for now)
func (a *Application) GetDBService() error {
	switch a.Config.Db.Driver {
	case Mysql:
		err := a.getMysqlServices()
		if err != nil {
			return err
		}
	case Inmemory:
		err := a.getInmemoryServices()
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown DB driver in config: %s", a.Config.Db.Driver)
	}

	return nil
}

func (a *Application) getMysqlServices() error {
	var err error
	a.DB, err = mysql.NewDB(a.Config.Db.DSN)
	if err != nil {
		return err
	}

	// Data services
	a.UserService = mysql.NewUserService(a.DB)
	a.TodoService = mysql.NewToDoService(a.DB)

	// Sessions storage (in app db, could be moved to Redis or some such if required for performance)
	a.SessionStore = mysqlstore.New(a.DB)

	return nil
}

func (a *Application) getInmemoryServices() error {
	userService := inmemory.NewUserService()
	todoService := inmemory.NewToDoService()

	if a.Config.Db.PreloadFile != "" {
		a.DebugMessage("Pre loading inmemory DB from file: %s\n", a.Config.Db.PreloadFile)

		err := userService.PreloadDataFromFile(a.Config.Db.PreloadFile)

		if err != nil {
			return fmt.Errorf("preloading inmemory user service error: %s", err)
		}

		err = todoService.PreloadDataFromFile(a.Config.Db.PreloadFile)

		if err != nil {
			return fmt.Errorf("preloading inmemory todo service error: %s", err)
		}
	}

	a.UserService = userService
	a.TodoService = todoService

	// Sessions storage
	a.SessionStore = memstore.New()

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

	// Do this so we can use ENV var's in the yaml
	for _, k := range viper.AllKeys() {
		v := viper.GetString(k)
		viper.Set(k, os.ExpandEnv(v))
	}

	hooks := mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
		HttpSamesiteFromStringViperHook(),
	)

	err = viper.Unmarshal(&config, viper.DecodeHook(hooks))
	if err != nil {
		return fmt.Errorf("cannot un-marshall config: %s", err)
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
