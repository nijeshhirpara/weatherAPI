package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendJsonResponse(w http.ResponseWriter, res interface{}) {
	jsonContent, jsonError := json.MarshalIndent(res, "", "	")
	if jsonError != nil {
		fmt.Println("HandleSettings: Error in JSON marshal", jsonError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonContent)
}

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
