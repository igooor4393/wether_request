package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Message struct {
	Items []data `json:"items"`
}

type data struct {
	ContentDetails float64 `json:"snippet"`
}

func main() {

	var m data

	url := "https://weatherbit-v1-mashape.p.rapidapi.com/current?lon=49.6668&lat=58.6035321&units=metric&lang=en"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "416d0e55acmsh81214bde525b3adp187c16jsn91d39d692e66")
	req.Header.Add("X-RapidAPI-Host", "weatherbit-v1-mashape.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	//defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	//fmt.Println(string(body))

	s := json.Unmarshal(body, &m.ContentDetails)

	fmt.Println(s)

}
