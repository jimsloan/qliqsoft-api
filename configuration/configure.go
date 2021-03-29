package configuration

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Config struct ...
type Config struct {
	Baseurl    string
	Email      string
	Token      string
	Report     string
	Filters    map[string]string
	Page       int
	PerPage    int
	Limit      int
	Outputpath string
}

// Configure ...
func Configure() Config {

	var apiconf Config
	var flags Cliflags
	var env Envars
	var cfg Yamlconfig

	flags = ParseFlags()
	env = Getenv()

	if flags.ConfigPath > "" {
		cfg = GetYaml(flags.ConfigPath)
	} else {
		log.Fatal("missing config file. Use flag -config\n")
	}

	apiconf.Email = env.AdminEmail
	apiconf.Token = env.Token
	apiconf.Baseurl = cfg.API.BaseURL

	if flags.Report > "" {
		fmt.Println("page:" + flags.Report)
		apiconf.Report = flags.Report
	} else {
		apiconf.Report = cfg.API.Endpoint
	}

	// merge filters from cli and yaml
	apiconf.Filters = mergefilters(flags.Filters, cfg.Filters)

	if flags.Page > 0 {
		fmt.Println("page:" + strconv.Itoa(flags.Page))
		apiconf.Page = flags.Page
	} else {
		apiconf.Page = cfg.Control.Page
	}

	if flags.PerPage > 0 {
		fmt.Println("perpage:" + strconv.Itoa(flags.PerPage))
		apiconf.PerPage = flags.PerPage
	} else {
		apiconf.PerPage = cfg.Control.PerPage
	}

	if flags.Limit > 0 {
		fmt.Println("limit:" + strconv.Itoa(flags.Limit))
		apiconf.Limit = flags.Limit
	} else {
		apiconf.Limit = cfg.Control.Limit
	}

	if flags.Outputpath > "" {
		fmt.Println("page:" + flags.Outputpath)
		apiconf.Outputpath = flags.Outputpath
	} else {
		apiconf.Outputpath = cfg.Control.Outputpath
	}

	return apiconf
}

// merge filters
func mergefilters(cli string, yaml map[string]interface{}) map[string]string {
	ret := make(map[string]string)

	// set the filters from the config file
	for fk, fv := range yaml {
		ret[fmt.Sprintf("%s", fk)] = fmt.Sprintf("%s", fv)
	}

	// update/add filters from cli
	if cli > "" {
		entries := strings.Split(cli, ",")
		for _, e := range entries {
			parts := strings.Split(e, "=")
			ret[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	return ret
}
