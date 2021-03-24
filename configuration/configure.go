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
	FromTime   string
	ToTime     string
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

	if flags.Filters > "" {

		// add parse for filters
		fmt.Println("filters:" + flags.Filters)
		s := strings.Split(flags.Filters, ",")
		fmt.Println(s)

		// use YAML fo now
		apiconf.FromTime = cfg.Filter.FromTime
		apiconf.ToTime = cfg.Filter.ToTime
	} else {
		apiconf.FromTime = cfg.Filter.FromTime
		apiconf.ToTime = cfg.Filter.ToTime
	}

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
