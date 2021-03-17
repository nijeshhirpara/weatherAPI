package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/stretchr/objx"
)

type weatherData struct {
	WindSpeed   float64 `json:"wind_speed"`
	Temperature float64 `json:"temperature_degrees"`
}

func HandleWeather(w http.ResponseWriter, r *http.Request) {
	var content Response
	content.Status = "success"

	city, ok := r.URL.Query()["city"]

	if !ok || len(city[0]) < 1 {
		log.Println("Url Param 'city' is missing")
		content.Status = "error"
		content.Data = "Url Param 'city' is missing"
		SendJsonResponse(w, content)
		return
	}

	if data, err := fetchFromWeatherStack(city[0]); err == nil {
		content.Data = data
		SendJsonResponse(w, content)
		return
	}

	SendJsonResponse(w, content)
	return
}

func fetchFromWeatherStack(city string) (weatherData, error) {
	var wd weatherData

	accessKey := os.Getenv("WeatherStack_Access_Key")
	url := fmt.Sprintf("http://api.weatherstack.com/current?access_key=%s&query=%s", accessKey, city)
	res, err := http.Get(url)
	if err != nil {
		return wd, err
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return wd, err
	}

	res.Body.Close()

	// Declared an empty interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal([]byte(resData), &result)

	obj := objx.New(result)
	if status := obj.Get("success").String(); status != "" && status == "false" {
		return wd, fmt.Errorf("Error fetching from weather stack API")
	}

	current := objx.New(result["current"])
	wd.Temperature = current.Get("temperature").Float64()
	wd.WindSpeed = current.Get("wind_speed").Float64()
	return wd, nil
}
