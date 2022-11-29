package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Cordinat struct {
	Data []struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		County    string  `json:"county"`
	} `json:"data"`
}

func Data(city string) (la, lo float64, ci string) {

	var LL Cordinat

	baseURL, _ := url.Parse("http://api.positionstack.com")

	baseURL.Path += "v1/forward"

	params := url.Values{}
	// Access Key
	params.Add("access_key", "f47e6a7ee827fbf3eee80c2eae6a4ef5")
	// Query
	params.Add("query", city)
	// Optional parameters
	//params.Add("region", "Washington")
	params.Add("output", "json")
	params.Add("limit", "1")

	baseURL.RawQuery = params.Encode()

	req, _ := http.NewRequest("GET", baseURL.String(), nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(string(body)) // ------ ПРОВЕРКА ТЕЛА ЗАПРОСА

	err := json.Unmarshal(body, &LL)
	if err != nil {
		log.Fatal("Unmarshal failed", err)
	}
	fmt.Println("Запрос координат")
	longitude := LL.Data[0].Longitude
	latitude := LL.Data[0].Latitude
	country := LL.Data[0].County

	//fmt.Println(longitude)
	//fmt.Println(latitude)
	//fmt.Println(country)

	return latitude, longitude, country
}
