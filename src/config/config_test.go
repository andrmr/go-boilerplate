package config

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	defaultTestPrefix = "TEST_APP_SETTINGS_"
)

var (
	defaultTestEnv = map[string]string{
		"TEST_APP_SETTINGS_NAME":               "John",
		"TEST_APP_SETTINGS_AGE":                "30",
		"TEST_APP_SETTINGS_DB_MANAGEMENT_HOST": "localhost1",
		"TEST_APP_SETTINGS_DB_MANAGEMENT_PORT": "123",
		"TEST_APP_SETTINGS_DB_STATISTICS_HOST": "localhost2",
		"TEST_APP_SETTINGS_DB_STATISTICS_PORT": "124",
	}

	defaultTestSettings = Settings{
		Name: "John",
		Age:  30,
		Db1: DbSettings{
			Host: "localhost1",
			Port: 123,
		},
		Db2: DbSettings{
			Host: "localhost2",
			Port: 124,
		},
	}
)

func TestDefaultEnv(t *testing.T) {
	for k, v := range defaultTestEnv {
		t.Setenv(k, v)
	}

	s, err := parse(defaultTestPrefix)
	if err != nil {
		t.Fatalf("parsing error: %v", err)
	}

	assert.Equal(t, defaultTestSettings, s, "objects should be equal")
}

func TestRequiredVar(t *testing.T) {
	testEnv := defaultTestEnv
	for k := range testEnv {
		delete(testEnv, k)
		break
	}

	for k, v := range testEnv {
		t.Setenv(k, v)
	}

	_, err := parse(defaultTestPrefix)
	assert.NotNil(t, err, "missing var should return error")
}

func TestRequiredType(t *testing.T) {
	testEnv := defaultTestEnv
	for k, v := range testEnv {
		if _, err := strconv.Atoi(v); err == nil {
			testEnv[k] = "zzz"
		}
	}

	for k, v := range testEnv {
		t.Setenv(k, v)
	}

	_, err := parse(defaultTestPrefix)
	assert.NotNil(t, err, "different type should return error")
}
