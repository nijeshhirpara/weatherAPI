package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

	var wd weatherData
	var err error

	wd, err = fetchFromWeatherStack(city[0])
	if err == nil {
		content.Data = wd
		SendJsonResponse(w, content)
		return
	}

	log.Println(err)

	wd, err = fetchFromOpenWeatherMap(city[0])
	if err == nil {
		content.Data = wd
		SendJsonResponse(w, content)
		return
	}

	log.Println(err)

	content.Status = "error"
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
		return wd, fmt.Errorf("Error fetching from weather stack API, Response %s", resData)
	}

	current := objx.New(result["current"])

	// Response temperature is in celsius
	wd.Temperature = math.Round((current.Get("temperature").Float64())*100) / 100

	// Response Wind Speed is in km/hr
	wd.WindSpeed = math.Round((current.Get("wind_speed").Float64())*100) / 100
	return wd, nil
}

func fetchFromOpenWeatherMap(city string) (weatherData, error) {
	var wd weatherData

	accessKey := os.Getenv("OpenWeatherMap_AppID")
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?appid=%s&q=%s", accessKey, city)
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
	if status := obj.Get("cod").String(); status != "200" {
		return wd, fmt.Errorf("Error fetching from Open weather Map API, Response %s", resData)
	}

	// Response temperature is in Kelvin, hence we need to convert it into Celsius
	main := objx.New(result["main"])
	wd.Temperature = math.Round((main.Get("temp").Float64()-273.15)*100) / 100 // Kelvin to Celsius

	// Response wind speed is in mt/sec, hence we need to convert it into km/hr
	wind := objx.New(result["wind"])
	wd.WindSpeed = math.Round((wind.Get("speed").Float64()*3.6)*100) / 100 // mt/sec to km/hr

	return wd, nil
}
