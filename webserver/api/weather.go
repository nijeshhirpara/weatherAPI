package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleWeather(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Data struct {
			WindSpeed   int `json:"wind_speed"`
			Temperature int `json:"temperature_degrees"`
		} `json:"data"`
		Status string `json:"status"`
	}

	var content Response
	content.Status = "success"
	content.Data.WindSpeed = 0
	content.Data.Temperature = 0

	jsonContent, jsonError := json.MarshalIndent(content, "", "	")
	if jsonError != nil {
		fmt.Println("HandleSettings: Error in JSON marshal", jsonError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonContent)
	return
}
