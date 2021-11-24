package configuration

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Cliflags struct ...
type Cliflags struct {
	ConfigPath    string
	Outputpath    string
	Report        string
	Filters       string
	Page          int
	PerPage       int
	Limit         int
	ClientTimeout int
}

// ParseFlags will create and parse the CLI flags
// and return the path to be used elsewhere
func ParseFlags() Cliflags {
	// Struct that contains the configured configuration path
	var flags Cliflags

	// CLI flag called "-config" to supply the configuration file
	flag.StringVar(&flags.ConfigPath, "config", "", "path to a config file")

	// CLI flag called "-outpath" to supply the path for the output
	flag.StringVar(&flags.Outputpath, "outpath", "", "path for the output")

	// CLI flag called "-report" to select the endpoint
	flag.StringVar(&flags.Report, "report", "", "report endpoint to query")

	// CLI flag called "-report" to select the endpoint
	flag.StringVar(&flags.Filters, "filters", "", "pass filters for the query")

	// CLI flag called "-page"
	flag.IntVar(&flags.Page, "page", 0, "starting page")

	// CLI flag called "-perpage"
	flag.IntVar(&flags.PerPage, "perpage", 0, "1 or more records per page")

	// CLI flag called "-limit"
	flag.IntVar(&flags.Limit, "limit", 0, "1 or more to limit the number of pages requested")

	// CLI flag called "-limit"
	flag.IntVar(&flags.ClientTimeout, "timeout", 0, "1 or more to set the client timeout, 30 seconds is the default")

	// Actually parse the flags
	flag.Parse()

	// Validate the path if set
	if flags.ConfigPath > "" {
		if err := ValidateConfigPath(flags.ConfigPath); err != nil {
			log.Fatal("invalid config path")
		}
	}

	// Validate the path if set
	if flags.Outputpath > "" {
		_, err := os.Stat(flags.Outputpath)
		if err != nil {
			log.Fatal("invalid output path")
		}
	}

	// Return the flags
	return flags
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
