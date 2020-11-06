package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
		Page        int64
		Items       int64
		Count       int64
		Total_pages int64
	}
}

func main() {

	url := "https://webprod.qliqsoft.com/quincy_api/v1/virtual_visits"
	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("from_time", "2020-11-05T10:00:00.000-06:00")
	q.Add("to_time", "2020-11-05T11:19:19.000-06:00")
	q.Add("page", "1")

	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", "{token}")
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

	// var result map[string]interface{}

	var result Response
	jsonErr := json.Unmarshal([]byte(body), &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// fmt.Printf("%s\n", string(body))
	// fmt.Printf("%+v\n", result)
	fmt.Printf("Page: %d\n", result.Meta.Page)
	fmt.Printf("Page Count: %d\n", result.Meta.Total_pages)
	for i, s := range result.Virtual_visits.Data {
		fmt.Printf("%d - %s ~ %v\n", i, s.Attributes.SessionID, s.Attributes.Duration)
	}
}
