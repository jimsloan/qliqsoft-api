package configuration

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config struct ...
type Config struct {
	API struct {
		BaseURL  string `yaml:"baseurl"`
		Endpoint string `yaml:"endpoint"`
	} `yaml:"api"`
	Filter struct {
		FromTime string `yaml:"fromtime"`
		ToTime   string `yaml:"totime"`
	}
	Control struct {
		Page    int `yaml:"startpage"`
		PerPage int `yaml:"recordsperpage"`
		Limit   int `yaml:"pagelimit"`
	} `yaml:"control"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
