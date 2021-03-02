package main

import (
	"github.com/jimsloan/qliqsoft-api/configuration"
	"github.com/jimsloan/qliqsoft-api/run"
)

func main() {
	run.Qliqsoft(configuration.Configure())
}
