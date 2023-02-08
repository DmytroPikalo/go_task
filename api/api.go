package api

import (
	"encoding/json"
	"net/http"
	"test_task/helpers"
)

var Data *helpers.MemoryData = new(helpers.MemoryData)

func Refresh(w http.ResponseWriter, r *http.Request) {
	var req helpers.JsonReader
	response := make(map[string]string)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// With goroutines, we don't have to wait for the completion
	go Data.ChangeApis(req)
	response["text"] = "The update has started"
	helpers.MakeJsonRes(w, response)
}

func LastChanges(w http.ResponseWriter, r *http.Request) {
	helpers.MakeJsonRes(w, Data.Changes)
}

func Filter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	if helpers.ValidateStr(query) {
		response := make(map[string][]string)
		response["match"] = Data.Match(query)
		helpers.MakeJsonRes(w, response)
	} else {
		response := make(map[string]string)
		response["error"] = "The string contains invalid characters"
		helpers.MakeJsonRes(w, response)
	}
}

func Count(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]int)
	response["count"] = len(Data.Ips)

	helpers.MakeJsonRes(w, response)
}

func Contains(w http.ResponseWriter, r *http.Request) {
	var req helpers.JsonReader
	response := make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if helpers.ValidateIp(req.IpAddress) {
		response["exists"] = Data.IpExists(req.IpAddress)
		helpers.MakeJsonRes(w, response)
	} else {
		response["error"] = "Ip not Valid"
		helpers.MakeJsonRes(w, response)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	Data.RLock()
	defer Data.RUnlock()
	Data.Ips = make(map[string]bool)
}
