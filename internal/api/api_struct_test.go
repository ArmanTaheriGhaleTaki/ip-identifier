package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_valid_ip(t *testing.T) {
	const input string = "12.12.41.15"
	const expected bool = true
	output := IpPatternValidation(input)
	assert.Equal(t, output, expected, "The two strings should be the same.")
}

func Test_invalid_ip(t *testing.T) {
	const input string = "258.12.41.15"
	const expected bool = false
	output := IpPatternValidation(input)
	assert.Equal(t, output, expected, "The two strings should be the same.")
}

// ############### it's not working cause of the connetion to database #########################
// func Test_ip_info_gather_info(t *testing.T) {
// 	const input string = "1.1.1.1"
// 	var expected = ip_struct.IpInfo{
// 		Ip:                      "1.1.1.1",
// 		Country:                 "Australia",
// 		CountryCode:             "AU",
// 		InternetServiceProvider: "Cloudflare, Inc",
// 	}
// 	output := IpInfoGatherInfo(input)
// 	assert.Equal(t, output, expected, "The two strings should be the same.")
// }
