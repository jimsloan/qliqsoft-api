package run

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jimsloan/qliqsoft-api/configuration"
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
func Qliqsoft(conf configuration.Config) {

	var result Response
	var pagecount int = 0

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
		err := WriteToJSON(conf.Outputpath, conf.Report, result.Meta.Page, data)
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
		if conf.Limit > 0 {
			pagecount++
			if pagecount == conf.Limit {
				break
			}
		}

		conf.Page++

	}
}

// WriteToJSON ...
func WriteToJSON(path string, name string, page int, data []byte) error {

	fmt.Println("writeToJSON ... page:" + strconv.Itoa(page))

	filename := path + name + "-" + strconv.Itoa(page) + ".json"
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
