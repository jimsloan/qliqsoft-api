package run

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/jimsloan/qliqsoft-api/fetch"
	"github.com/jimsloan/qliqsoft-api/output"
)

var conf fetch.Config

// Data ...
type Data struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Attributes map[string]interface{} `json:"attributes"`
}

// Response ...
type Response struct {
	VirtualVisits struct {
		Data []Data
	} `json:"virtual_visits"`
	Meta struct {
		Filters struct {
			FromTime string `json:"from_time"`
			ToTime   string `json:"to_time"`
		}
		Page       int `json:"page"`
		Items      int `json:"items"`
		Count      int `json:"count"`
		TotalPages int `json:"total_pages"`
	}
}

// Qliqsoft ...
func Qliqsoft() {

	var result Response
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

	doJSON := false
	doCSV := true

	// call fetchAPI() until there are no more pages
	var arrayrow [][]string

	for {

		// fetch API data
		data := fetch.API(conf)

		// move json to structs
		jsonErr := json.Unmarshal([]byte(data), &result)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		if doJSON {
			err := output.WriteToJSON(result.Meta.Page, data)
			if err != nil {
				log.Fatal(err)
			}
		}

		// output results
		if result.Meta.Count > 0 {
			fmt.Printf("Fetching page %d of %d; %d items (%d)\n", result.Meta.Page, result.Meta.TotalPages, result.Meta.Items, result.Meta.Count)
		} else {
			fmt.Printf("No records returned.\n")
			break
		}

		// sanity check: page count should never exceed the total pages
		if conf.Page > result.Meta.TotalPages {
			log.Fatal("Page count exceeded total pages")
		}

		if doCSV {
			var row []string

			// create header from sorted map keys
			keys := make([]string, 0, len(result.VirtualVisits.Data[0].Attributes))
			for k := range result.VirtualVisits.Data[0].Attributes {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			if result.Meta.Page == 1 {
				arrayrow = append(arrayrow, keys)
			}

			for _, s := range result.VirtualVisits.Data {
				for _, k := range keys {
					row = append(row, fmt.Sprint(s.Attributes[k]))
				}
				arrayrow = append(arrayrow, row)
				row = nil
			}

			//writeToCsv(result.Meta.Page, result)
		}

		// check page count and either repeat or exit
		if conf.Page == result.Meta.TotalPages {
			break
		}

		conf.Page++

	}
	if doCSV {
		output.WriteToCsv(arrayrow)
	}
}
