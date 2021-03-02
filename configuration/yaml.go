package configuration

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Yamlconfig struct ...
type Yamlconfig struct {
	API struct {
		BaseURL  string `yaml:"baseurl"`
		Endpoint string `yaml:"endpoint"`
	} `yaml:"api"`
	Filter struct {
		FromTime string `yaml:"fromtime"`
		ToTime   string `yaml:"totime"`
	}
	Control struct {
		Page       int    `yaml:"startpage"`
		PerPage    int    `yaml:"recordsperpage"`
		Limit      int    `yaml:"pagelimit"`
		Outputpath string `yaml:"outputpath"`
	} `yaml:"control"`
}

// GetYaml returns a new decoded Config struct
func GetYaml(configPath string) Yamlconfig {
	// Create config structure
	var config Yamlconfig

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		log.Fatal(err)
	}

	return config
}
