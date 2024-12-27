package api

import (
	"encoding/json"
	"io"
	"log"
	"mmd/internal/crud"
	ip_struct "mmd/internal/ip"
	"net/http"
)

const (
	ip_query_url string = "http://ip-api.com/json/"
)

func IP_info_gather_info(ip string) ip_struct.IP_info {
	var result ip_struct.IP_info
	result.Ip = ip
	crud.Database_retrieve(&result)
	if result.Country == "" {
		url := ip_query_url + ip
		log.Printf("lookup by the remote api: %v", ip)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("failed to connect to url: %v", err)
		}
		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("failed to read response body: %v", err)
		}
		err = json.Unmarshal(bodyBytes, &result)
		if err != nil {
			log.Fatalf("failed to unmarshal response body: %v", err)
		}
		crud.Database_insert(&result)
		result.Ip = ip
	}
	return result
}
