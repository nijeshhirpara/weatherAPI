package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendJsonResponse converts interface into json and writes an output with json header
func SendJsonResponse(w http.ResponseWriter, res interface{}) {
	jsonContent, jsonError := json.MarshalIndent(res, "", "	")
	if jsonError != nil {
		fmt.Println("HandleSettings: Error in JSON marshal", jsonError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonContent)
}

// Response is a reference for sending output
type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
