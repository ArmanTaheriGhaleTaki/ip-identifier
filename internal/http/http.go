package api

import (
	"encoding/json"
	"log"
	internal "mmd/internal/api"
	"net/http"
)

var api_path string = "/info"

func Http_start() {

	http.HandleFunc(api_path, helloHandler)
	log.Println("Starting web server liten on port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	if internal.Ip_pattern_validation(ip) {
		ip_info := internal.IP_info_gather_info(ip)
		if err := json.NewEncoder(w).Encode(ip_info); err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

}
