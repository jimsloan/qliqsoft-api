package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Attributes struct {
	CallType      string
	Owner         string
	Widget        string
	SessionID     string
	Type          string
	Title         string
	Status        string
	StartAt       string
	JoinedAt      string
	LeftAt        string
	Duration      float64
	DeviceBrowser string
	FailureReason string
}

type Data struct {
	Id         string
	Type       string
	Attributes Attributes
}

type Response struct {
	Virtual_visits struct {
		Data []Data
	}
	Meta struct {
		Filters struct {
			From_time string
			To_time   string
		}
		Page        int
		Items       int
		Count       int
		Total_pages int
	}
}

func fetchAPI(url string, token string, from string, to string, page int) []byte {
	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("from_time", from)
	q.Add("to_time", to)
	q.Add("page", strconv.Itoa(page))

	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", token)
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

func main() {

	token, ok := os.LookupEnv("QLIQ_API_TOKEN")
	if !ok {
		log.Fatal("QLIQ_API_TOKEN not set\n")
	}
	if len(token) == 0 {
		log.Fatal("QLIQ_API_TOKEN empty\n")
	}

	url := "https://webprod.qliqsoft.com/quincy_api/v1/virtual_visits"
	fromTime := "2020-11-05T00:00:00.000-06:00"
	toTime := "2020-11-05T06:00:00.000-06:00"

	// var result map[string]interface{}
	var result Response

	// call fetchAPI() until there are no more pages
	totalPages := 1

	for thisPage := 0; thisPage < totalPages; thisPage++ {
		data := fetchAPI(url, token, fromTime, toTime, thisPage+1)

		jsonErr := json.Unmarshal([]byte(data), &result)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		totalPages = result.Meta.Total_pages

		// loop over the visit records
		for i, s := range result.Virtual_visits.Data {
			fmt.Printf("%d - %s ~ %v\n", i, s.Attributes.SessionID, s.Attributes.Duration)
		}

		if thisPage+1 != result.Meta.Page {
			fmt.Printf("Pages out of sync.\n")
			break
		}

		if result.Meta.Count > 0 {
			fmt.Printf("Page %d of %d; %d items (%d)\n", result.Meta.Page, result.Meta.Total_pages, result.Meta.Items, result.Meta.Count)
			fmt.Printf("------\n")
		} else {
			fmt.Printf("No records returned.\n")
			break
		}

	}
}
