package ip

type IP_info struct {
	Ip                        string `json:"ip" gorm:"primaryKey"`
	Country                   string `json:"country"`
	Country_code              string `json:"countryCode"`
	Internet_service_provider string `json:"isp"`
}
