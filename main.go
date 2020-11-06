package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Record struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes map[string]interface{}
}

type Visits struct {
	virtual_visits struct {
		Data []Record
	}
	Meta struct {
		Filters struct {
			From string `json:"from_time"`
			To   string `json:"to_time"`
		}
		Page      int64 `json:"page"`
		Items     int64 `json:"items"`
		Count     int64 `json:"count"`
		PageCount int64 `json:"total_pages"`
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
	q.Add("from_time", "2020-11-05T11:19:19.000-06:00")
	q.Add("to_time", "2020-11-05T11:19:19.000-06:00")
	q.Add("page", "1")

	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", "{insert toked here}")
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

	var result Visits
	//var result map[string]interface{}
	jsonErr := json.Unmarshal([]byte(body), &result)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Printf("%+v\n", result)
	// fmt.Printf("Page: %d\n", result.Meta.Page)
	// fmt.Printf("Page Count: %d\n", result.Meta.PageCount)

}
