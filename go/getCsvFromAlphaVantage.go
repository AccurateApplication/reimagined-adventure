package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

func ReadCSVfromURL(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	url := "" // Url with api key
	data, err := ReadCSVfromURL(url)
	if err != nil {
		panic(err)
	}
	for idx, row := range data {
		//		if idx == 0 {
		//			continue
		//		}
		if idx == 3 {
			break
		}
		fmt.Println(row)
	}
	fmt.Printf("type of data %T: \n", data)
}
