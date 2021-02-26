package fetch

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Config ...
type Config struct {
	Token    string
	Email    string
	URL      string
	Endpoint string
	FromTime string
	ToTime   string
	Page     int
	PerPage  int
}

//API ...
func API(runtime Config) []byte {
	client := http.Client{
		Timeout: time.Second * 30, // Timeout after 2 seconds
	}

	req, err := http.NewRequest("GET", runtime.URL+runtime.Endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	if runtime.FromTime > "" && runtime.ToTime > "" {
		q.Add("from_time", runtime.FromTime)
		q.Add("to_time", runtime.ToTime)
	}
	q.Add("page", strconv.Itoa(runtime.Page))
	q.Add("per_page", strconv.Itoa(runtime.PerPage))

	req.URL.RawQuery = q.Encode()

	req.SetBasicAuth(runtime.Email, runtime.Token)
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
