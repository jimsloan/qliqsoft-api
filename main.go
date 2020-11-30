package main

import (
	"log"
	"os"

	"github.com/jimsloan/qliqsoft-api/fetch"
	"github.com/jimsloan/qliqsoft-api/run"
)

func main() {

	var conf fetch.Config

	// pass secrets via environment
	token, ok := os.LookupEnv("QLIQ_API_TOKEN")
	if !ok {
		log.Fatal("QLIQ_API_TOKEN not set\n")
	}
	if len(token) == 0 {
		log.Fatal("QLIQ_API_TOKEN empty\n")
	}
	conf.Token = token

	// need to move these to parameters or config file
	conf.URL = "https://webprod.qliqsoft.com/quincy_api/v1/virtual_visits"
	conf.FromTime = "2020-11-05T00:00:00.000-06:00"
	conf.ToTime = "2020-11-05T09:00:00.000-06:00"
	conf.Page = 1
	run.Run(conf)
}
