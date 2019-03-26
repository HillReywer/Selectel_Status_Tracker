package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://selectel.status.io/1.0/status/5980813dd537a2a7050004bd"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}

func UnmarshalSelectel(data []byte) (Selectel, error) {
	var r Selectel
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Selectel) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Selectel struct {
	Result Result `json:"result"`
}

type Result struct {
	StatusOverall StatusOverall   `json:"status_overall"`
	Status        []StatusElement `json:"status"`
	Incidents     []interface{}   `json:"incidents"`
	Maintenance   Maintenance     `json:"maintenance"`
}

type Maintenance struct {
	Active   []interface{} `json:"active"`
	Upcoming []interface{} `json:"upcoming"`
}

type StatusElement struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	Updated    string          `json:"updated"`
	Status     StatusEnum      `json:"status"`
	StatusCode int64           `json:"status_code"`
	Containers []StatusElement `json:"containers"`
}

type StatusOverall struct {
	Updated    string     `json:"updated"`
	Status     StatusEnum `json:"status"`
	StatusCode int64      `json:"status_code"`
}

type StatusEnum string
const (
	Operational StatusEnum = "Operational"
)
