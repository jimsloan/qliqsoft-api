package run

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jimsloan/qliqsoft-api/fetch"
)

// Response ...
type Response struct {
	Meta struct {
		Page       int `json:"page"`
		Items      int `json:"items"`
		Count      int `json:"count"`
		TotalPages int `json:"total_pages"`
	} `json:"meta"`
}

// Qliqsoft ...
func Qliqsoft(endpoint string) {

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
	conf.URL = "https://webprod.qliqsoft.com/quincy_api/v1/reports/"
	conf.Endpoint = endpoint
	conf.FromTime = "" // "2021-02-24T17:00:00-06:00"
	conf.ToTime = ""   //"2021-02-25T16:59:59-06:00"
	conf.Page = 1
	conf.PerPage = 1
	conf.Email = "hagedol9@trinity-health.org"

	limitPages := 1

	// call fetchAPI() until there are no more pages
	for {

		// fetch API data
		data := fetch.API(conf)

		// move json to structs
		jsonErr := json.Unmarshal([]byte(data), &result)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		// write out the json
		err := WriteToJSON(endpoint, result.Meta.Page, data)
		if err != nil {
			log.Fatal(err)
		}

		// send fome logging info to stdout
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

		// check page count and either continue or exit
		if conf.Page == result.Meta.TotalPages {
			break
		}

		// is there a page limit on the requests
		if limitPages > 0 {
			if conf.Page == limitPages {
				break
			}
		}

		conf.Page++

	}
}

// WriteToJSON ...
func WriteToJSON(name string, page int, data []byte) error {

	fmt.Println("writeToJSON ... page:" + strconv.Itoa(page))

	filename := "./var/" + name + "-" + strconv.Itoa(page) + ".json"
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return file.Sync()
}
