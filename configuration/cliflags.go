package configuration

import (
	"flag"
	"fmt"
	"os"
)

// CliFlag struct ...
type CliFlag struct {
	ConfigPath string
	Limit      int
}

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags() (CliFlag, error) {
	// Struct that contains the configured configuration path
	var flags CliFlag

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&flags.ConfigPath, "config", "./config.ymal", "path to config file")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	if err := ValidateConfigPath(flags.ConfigPath); err != nil {
		return flags, err
	}

	// Return the flags
	return flags, nil
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}
