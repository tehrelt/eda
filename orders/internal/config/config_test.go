package config_test

import (
	"os"
	"testing"

	"github.com/tehrelt/eda/orders/internal/config"
)

func assertPanic(t *testing.T, f func()) {

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	f()
}

func TestEmptyEnv(t *testing.T) {
	assertPanic(t, func() {
		config.New()
	})

}

func TestNew(t *testing.T) {
	envs := map[string]string{
		"APP_PORT":      "8080",
		"APP_ENV":       "prod",
		"DATABASE_USER": "user",
		"DATABASE_PASS": "pass",
		"DATABASE_NAME": "orders",
		"DATABASE_HOST": "localhost",
		"DATABASE_PORT": "5432",
	}

	for k, v := range envs {
		os.Setenv(k, v)
	}

	got := config.New()

	if got.App.Port != 8080 {
		t.Errorf("App.Port = %d, want %d", got.App.Port, 8080)
	}

	if got.App.Env != "prod" {
		t.Errorf("App.Env = %s, want %s", got.App.Env, "prod")
	}

	if got.Database.Host != "localhost" {
		t.Errorf("Database.Host = %s, want %s", got.Database.Host, "localhost")
	}

	if got.Database.Port != 5432 {
		t.Errorf("Database.Port = %d, want %d", got.Database.Port, 5432)
	}

	if got.Database.Name != "orders" {
		t.Errorf("Database.Name = %s, want %s", got.Database.Name, "orders")
	}

	if got.Database.User != "user" {
		t.Errorf("Database.User = %s, want %s", got.Database.User, "user")
	}

	if got.Database.Pass != "pass" {
		t.Errorf("Database.Pass = %s, want %s", got.Database.Pass, "pass")
	}
}
