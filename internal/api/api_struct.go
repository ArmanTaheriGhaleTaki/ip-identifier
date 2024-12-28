package api

import (
	"IpIdentifier/internal/crud"
	ip_struct "IpIdentifier/internal/ip"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const (
	IPQEURYURL string = "http://ip-api.com/json/"
)

func IpInfoGatherInfo(ip string) ip_struct.IpInfo {
	var result ip_struct.IpInfo
	result.Ip = ip
	crud.DatabaseRetrieve(&result)
	if result.Country == "" {
		url := IPQEURYURL + ip
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
		crud.DatabaseInsert(&result)
		result.Ip = ip
	}
	return result
}
