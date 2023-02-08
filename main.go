package main

import (
	"log"
	"net/http"

	"test_task/api"
)

func main() {
	// Read url from config file
	api.Data.GetConf()
	// read and save data to memory
	api.Data.LinesToMap()

	http.HandleFunc("/refresh", api.Refresh)
	http.HandleFunc("/last-changes", api.LastChanges)
	http.HandleFunc("/filter", api.Filter)
	http.HandleFunc("/count", api.Count)
	http.HandleFunc("/contains", api.Contains)
	http.HandleFunc("/delete", api.Delete)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println("Something going wrong: ", err)
		return
	}
}
