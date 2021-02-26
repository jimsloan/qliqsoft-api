package main

import (
	"github.com/jimsloan/qliqsoft-api/run"
)

func main() {
	// API endpoints
	// patients, agents, widgets, conversations, invitation_history, terms_and_conditions, opt_outs, forms, device_tests, virtual_visits

	run.Qliqsoft("widgets")
}
