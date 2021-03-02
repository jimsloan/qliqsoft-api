package configuration

import "log"

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
		apiconf.Report = flags.Report
	} else {
		apiconf.Report = cfg.API.Endpoint
	}

	if flags.FromTime > "" && flags.ToTime > "" {
		apiconf.FromTime = flags.FromTime
		apiconf.ToTime = flags.ToTime
	} else {
		apiconf.FromTime = cfg.Filter.FromTime
		apiconf.ToTime = cfg.Filter.ToTime
	}

	if flags.Page > 0 {
		apiconf.Page = flags.Page
	} else {
		apiconf.Page = cfg.Control.Page
	}

	if flags.PerPage > 0 {
		apiconf.PerPage = flags.PerPage
	} else {
		apiconf.PerPage = cfg.Control.PerPage
	}

	if flags.Limit > 0 {
		apiconf.Limit = flags.Limit
	} else {
		apiconf.Limit = cfg.Control.Limit
	}

	if flags.Outputpath > "" {
		apiconf.Outputpath = flags.Outputpath
	} else {
		apiconf.Outputpath = cfg.Control.Outputpath
	}

	return apiconf
}
