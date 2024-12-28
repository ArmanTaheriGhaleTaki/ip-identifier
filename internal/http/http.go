package api

import (
	"encoding/json"
	"log"
	internal "mmd/internal/api"
	"net/http"
)

const APIPATH string = "/info"

func HttpStart() {
	http.HandleFunc(APIPATH, helloHandler)
	log.Println("Starting web server liten on port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	if internal.IpPatternValidation(ip) {
		ipInfo := internal.IpInfoGatherInfo(ip)
		if err := json.NewEncoder(w).Encode(ipInfo); err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "bad request", http.StatusBadRequest)
	}
}
