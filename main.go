package main

import (
	"log"

	"github.com/jimsloan/qliqsoft-api/configuration"
	"github.com/jimsloan/qliqsoft-api/fetch"
	"github.com/jimsloan/qliqsoft-api/run"
)

func main() {

	// Generate our config based on the config supplied
	// by the user in the flags
	flags, err := configuration.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := configuration.NewConfig(flags.ConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	secrets := configuration.GetSecrets()

	var limit int = cfg.Control.Limit
	var apiconf fetch.Config

	apiconf.Email = secrets.AdminEmail
	apiconf.Token = secrets.Token
	apiconf.URL = cfg.API.BaseURL
	apiconf.Endpoint = cfg.API.Endpoint
	apiconf.FromTime = cfg.Filter.FromTime
	apiconf.ToTime = cfg.Filter.ToTime
	apiconf.Page = cfg.Control.Page
	apiconf.PerPage = cfg.Control.PerPage

	run.Qliqsoft(apiconf, limit)
}
