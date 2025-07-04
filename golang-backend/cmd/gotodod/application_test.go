package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfigPath(t *testing.T) {
	app := NewApplication()
	app.LoadConfig([]string{})

	expected := "/etc/gotodod"

	if app.ConfigPath != expected {
		t.Errorf("config path wrong, expected \"%s\", got \"%s\"", expected, app.ConfigPath)
	}
}

func TestExpandConfigPath(t *testing.T) {
	app := NewApplication()
	app.LoadConfig([]string{"-config", "~/.gotodod"})

	dirname, _ := os.UserHomeDir()
	expected := filepath.Join(dirname, ".gotodod")

	if app.ConfigPath != expected {
		t.Errorf("config path wrong, expected \"%s\", got \"%s\"", expected, app.ConfigPath)
	}
}

func TestDifferentConfigPath(t *testing.T) {
	expected := "/var/etc/gotodod"

	app := NewApplication()
	app.LoadConfig([]string{"-config", "/var/etc/gotodod"})

	if app.ConfigPath != expected {
		t.Errorf("config path wrong, expected \"%s\", got \"%s\"", expected, app.ConfigPath)
	}
}
