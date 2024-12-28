package ip

type IpInfo struct {
	Ip                      string `json:"ip" gorm:"primaryKey"`
	Country                 string `json:"country"`
	CountryCode             string `json:"countryCode"`
	InternetServiceProvider string `json:"isp"`
}
