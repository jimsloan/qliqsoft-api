package configuration

import (
	"flag"
	"fmt"
	"os"
)

// CliFlag struct ...
type CliFlag struct {
	ConfigPath string
	Report     string
	FromTime   string
	ToTime     string
	Page       int
	PerPage    int
	Limit      int
}

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags() (CliFlag, error) {
	// Struct that contains the configured configuration path
	var flags CliFlag

	// CLI flag called "-config" to supply the configuration file
	flag.StringVar(&flags.ConfigPath, "config", "", "path to config file")

	// CLI flag called "-report" to select the endpoint
	flag.StringVar(&flags.Report, "report", "", "report endpoint to query")

	// CLI flag called "-page"
	flag.IntVar(&flags.Page, "page", 0, "1 or more starting page")

	// CLI flag called "-perpage"
	flag.IntVar(&flags.PerPage, "perpage", 0, "1 or more records per page")

	// CLI flag called "-limit"
	flag.IntVar(&flags.Limit, "limit", 0, "1 or more to limit the number of pages requested")

	// Actually parse the flags
	flag.Parse()

	// Validate the path if set
	if flags.ConfigPath > "" {
		if err := ValidateConfigPath(flags.ConfigPath); err != nil {
			return flags, err
		}
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
