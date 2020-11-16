package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"
)

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

// Runtime ...
type Runtime struct {
	token    string
	url      string
	fromTime string
	toTime   string
	page     int
}

func run(runtime Runtime) {

	var result Response

	// call fetchAPI() until there are no more pages
	for {

		// request data
		data := fetchAPI(runtime)

		jsonErr := json.Unmarshal([]byte(data), &result)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		// output results
		if result.Meta.Count > 0 {
			fmt.Printf("Page %d of %d; %d items (%d)\n", result.Meta.Page, result.Meta.TotalPages, result.Meta.Items, result.Meta.Count)
		} else {
			fmt.Printf("No records returned.\n")
			break
		}

		doJSON := false
		if doJSON {
			err := writeToJSON(result.Meta.Page, data)
			if err != nil {
				log.Fatal(err)
			}
		}

		doCSV := false
		if doCSV {
			writeToCsv(result.Meta.Page, result)
		}

		// sanity check: page count should never exceed the total pages
		if runtime.page > result.Meta.TotalPages {
			log.Fatal("Page count exceeded total pages")
		}

		// check page count and either repeat or exit
		if runtime.page == result.Meta.TotalPages {
			break
		}
		runtime.page++

	}
}

func fetchAPI(runtime Runtime) []byte {
	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest("GET", runtime.url, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("from_time", runtime.fromTime)
	q.Add("to_time", runtime.toTime)
	q.Add("page", strconv.Itoa(runtime.page))

	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", runtime.token)
	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}

func writeToJSON(page int, data []byte) error {

	fmt.Println("writeToJSON ... page:" + strconv.Itoa(page))

	filename := "page-" + strconv.Itoa(page) + ".json"
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

func writeToCsv(page int, result Response) {

	// create header from sorted map keys
	keys := make([]string, 0, len(result.VirtualVisits.Data[0].Attributes))
	for k := range result.VirtualVisits.Data[0].Attributes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	path := "./data.csv"

	if page == 1 {
		// remove file if it exist
		if _, err := os.Stat(path); err == nil {
			err := os.Remove(path)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	fmt.Println("writeToCsv ... page:" + strconv.Itoa(page))

	csvFile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	writer := csv.NewWriter(csvFile)

	// write the headers
	if page == 1 {
		fmt.Println("writeToCsv ... Header")
		writer.Write(keys)
	}

	for _, s := range result.VirtualVisits.Data {
		var row []string
		for _, k := range keys {
			row = append(row, fmt.Sprint(s.Attributes[k]))
		}
		writer.Write(row)
	}
	writer.Flush()
}

func main() {

	var runtime Runtime

	// pass secrets via environment
	token, ok := os.LookupEnv("QLIQ_API_TOKEN")
	if !ok {
		log.Fatal("QLIQ_API_TOKEN not set\n")
	}
	if len(token) == 0 {
		log.Fatal("QLIQ_API_TOKEN empty\n")
	}

	runtime.token = token

	// need to move these to parameters or config file
	runtime.url = "https://webprod.qliqsoft.com/quincy_api/v1/virtual_visits"
	runtime.fromTime = "2020-11-05T00:00:00.000-06:00"
	runtime.toTime = "2020-11-05T09:00:00.000-06:00"
	runtime.page = 1
	run(runtime)
}
