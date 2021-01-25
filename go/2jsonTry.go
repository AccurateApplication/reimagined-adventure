package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// https://www.alphavantage.co/documentation/

func main() {
	//resp, _ := http.Get("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=IBM&apikey=XXX")
	resp, _ := http.Get("") // alpha vantage url with api
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// make map that we can decode json to
	result := make(map[string]interface{})

	if err := json.Unmarshal(byt, &result); err != nil {
		panic(err)
	}
	//fmt.Println(result)
	for _, i := range result {
		fmt.Println("i: ", i)
	}
}
