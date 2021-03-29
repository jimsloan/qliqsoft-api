package fetch

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jimsloan/qliqsoft-api/configuration"
)

//API ...
func API(runtime configuration.Config) []byte {

	// create the http client
	client := http.Client{
		Timeout: time.Second * 30, // Timeout after 2 seconds
	}

	// init the request
	req, err := http.NewRequest("GET", runtime.Baseurl+runtime.Report, nil)
	if err != nil {
		log.Fatal(err)
	}

	// init the query string
	q := req.URL.Query()

	// add the page controls
	q.Add("page", strconv.Itoa(runtime.Page))
	q.Add("per_page", strconv.Itoa(runtime.PerPage))

	// add the filters

	fmt.Print("\n--[Filters]--\n")
	fmt.Printf("%#+v", runtime.Filters)
	// for a, b := range runtime.Filters {
	// 	fmt.Printf("\t%s = %s\n", a, b)
	// }

	//if runtime.FromTime > "" && runtime.ToTime > "" {
	q.Add("from_time", "2021-03-28T00:00:00-06:00")
	//q.Add("to_time", "2021-03-29T00:00:00-06:00")
	//}

	req.URL.RawQuery = q.Encode()

	// basic auth credentials
	req.SetBasicAuth(runtime.Email, runtime.Token)

	// do the request
	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// read the response
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}
