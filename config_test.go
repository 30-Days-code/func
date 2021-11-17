//go:build !integration
// +build !integration

package function_test

// Config tests do not have private access, as they are testing manifest
// behavior of the public package interface.  For implementation tests, see
// files suffixed '_unit_test.go'.

import (
	"testing"
)

// TestConfigPathDefault ensures that config defaults to XDG_CONFIG_HOME/func
func TestConfigPathDefault(t *testing.T) {
	// TODO
	// Set XDG_CONFIG_PATH to ./testdata/config
	// Confirm the config is populated from the test files.
}

// TestConfigPath ensure that the config path provided via the WithConfig
// option is respected.
func TestConfigPath(t *testing.T) {
	// TODO
	// Create a client specifying ./testdata/config
	// Confirm the config is populated from the test files.
}

// TestConfigRepositoriesPath ensures that the repositories directory within
// the effective config path is created if it does not already exist.
func TestConfigRepositoriesPath(t *testing.T) {
	// TODO
	// Create a temporary directory
	// Specify this directory as the config path when instantiating a client.
	// Confirm that the repositories directory is created.
}
