package main

import (
	"net/http"
	"testing"

	"github.com/spf13/viper"
)

func TestSessionSameSite1(t *testing.T) {
	viper.Reset()
	app := NewApplication()
	t.Setenv("APP_ENV", "samesite1")
	app.LoadConfig([]string{"-config", "./test-configs"})

	expected := http.SameSiteStrictMode

	if app.Config.Session.SameSite != expected {
		t.Errorf("session same site wrong, expected \"%d\", got \"%d\"", expected, app.Config.Session.SameSite)
	}
}

func TestSessionSameSite2(t *testing.T) {
	viper.Reset()
	app := NewApplication()
	t.Setenv("APP_ENV", "samesite2")
	app.LoadConfig([]string{"-config", "./test-configs"})

	expected := http.SameSiteDefaultMode

	if app.Config.Session.SameSite != expected {
		t.Errorf("session same site wrong, expected \"%d\", got \"%d\"", expected, app.Config.Session.SameSite)
	}
}

func TestSessionSameSite3(t *testing.T) {
	viper.Reset()
	app := NewApplication()
	t.Setenv("APP_ENV", "samesite3")
	app.LoadConfig([]string{"-config", "./test-configs"})

	expected := http.SameSiteLaxMode

	if app.Config.Session.SameSite != expected {
		t.Errorf("session same site wrong, expected \"%d\", got \"%d\"", expected, app.Config.Session.SameSite)
	}
}

func TestSessionSameSite4(t *testing.T) {
	viper.Reset()
	app := NewApplication()
	t.Setenv("APP_ENV", "samesite4")
	app.LoadConfig([]string{"-config", "./test-configs"})

	expected := http.SameSiteNoneMode

	if app.Config.Session.SameSite != expected {
		t.Errorf("session same site wrong, expected \"%d\", got \"%d\"", expected, app.Config.Session.SameSite)
	}
}

// samesite set to nonsense so should default to strict
func TestSessionSameSite5(t *testing.T) {
	viper.Reset()
	app := NewApplication()
	t.Setenv("APP_ENV", "samesite5")
	app.LoadConfig([]string{"-config", "./test-configs"})

	expected := http.SameSiteStrictMode

	if app.Config.Session.SameSite != expected {
		t.Errorf("session same site wrong, expected \"%d\", got \"%d\"", expected, app.Config.Session.SameSite)
	}
}

func TestUnknownDbDriverOption(t *testing.T) {
	viper.Reset()
	app := NewApplication()
	t.Setenv("APP_ENV", "unknownDbDriver")
	app.LoadConfig([]string{"-config", "./test-configs"})

	err := app.GetDBService()
	if err == nil {
		t.Errorf("expected an error from unknown db driver spec")
	}
}
