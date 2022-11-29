package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Wether struct {
	Count int `json:"count"`
	Data  []struct {
		AppTemp     float64  `json:"app_temp"`
		Aqi         int      `json:"aqi"`
		CityName    string   `json:"city_name"`
		Clouds      int      `json:"clouds"`
		CountryCode string   `json:"country_code"`
		Datetime    string   `json:"datetime"`
		Dewpt       float64  `json:"dewpt"`
		Dhi         float64  `json:"dhi"`
		Dni         float64  `json:"dni"`
		ElevAngle   float64  `json:"elev_angle"`
		Ghi         float64  `json:"ghi"`
		Gust        float64  `json:"gust"`
		HAngle      int      `json:"h_angle"`
		Lat         float64  `json:"lat"`
		Lon         float64  `json:"lon"`
		ObTime      string   `json:"ob_time"`
		Pod         string   `json:"pod"`
		Pres        float64  `json:"pres"`
		Rh          int      `json:"rh"`
		Slp         float64  `json:"slp"`
		Snow        int      `json:"snow"`
		SolarRad    float64  `json:"solar_rad"`
		Sources     []string `json:"sources"`
		StateCode   string   `json:"state_code"`
		Station     string   `json:"station"`
		Sunrise     string   `json:"sunrise"`
		Sunset      string   `json:"sunset"`
		Temp        float64  `json:"temp"`
		Timezone    string   `json:"timezone"`
		Ts          int      `json:"ts"`
		Uv          float64  `json:"uv"`
		Vis         int      `json:"vis"`
		Weather     struct {
			Description string `json:"description"`
			Code        int    `json:"code"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		WindCdir     string  `json:"wind_cdir"`
		WindCdirFull string  `json:"wind_cdir_full"`
		WindDir      int     `json:"wind_dir"`
		WindSpd      float64 `json:"wind_spd"`
	} `json:"data"`
}

func Wetherr(la, lo float64) (temp float64) {

	var i Wether
	var s float64

	x := fmt.Sprintf("%.1f", la)
	y := fmt.Sprintf("%.1f", lo)

	url := fmt.Sprintf("https://weatherbit-v1-mashape.p.rapidapi.com/current?lon=%s&lat=%s", x, y)
	url = strings.Trim(url, "\n")

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "416d0e55acmsh81214bde525b3adp187c16jsn91d39d692e66")
	req.Header.Add("X-RapidAPI-Host", "weatherbit-v1-mashape.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err := json.Unmarshal(body, &i)
	if err != nil {
		log.Fatal("Unmarshal failed", err)
	}
	fmt.Println("Запрос погоды")
	fmt.Println(i.Count)
	s = i.Data[0].Temp
	fmt.Println(s)

	// ________________________________________________________

	return s
}
