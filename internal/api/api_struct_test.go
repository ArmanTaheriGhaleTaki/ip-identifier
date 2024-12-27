package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_valid_ip(t *testing.T) {
	const input string = "12.12.41.15"
	const expected bool = true
	output := Ip_pattern_validation(input)

	assert.Equal(t, output, expected, "The two strings should be the same.")

}
func Test_invalid_ip(t *testing.T) {
	const input string = "258.12.41.15"
	const expected bool = false
	output := Ip_pattern_validation(input)

	assert.Equal(t, output, expected, "The two strings should be the same.")

}

//############### it's not working cause of the connetion to database #########################
// func Test_ip_info_gather_info(t *testing.T) {
// 	const input string = "1.1.1.1"
// 	var expected = ip.IP_info{
// 		Ip:                        "1.1.1.1",
// 		Country:                   "Australia",
// 		Country_code:              "AU",
// 		Internet_service_provider: "Cloudflare, Inc",
// 	}
// 	output := IP_info_gather_info(input)

// 	assert.Equal(t, output, expected, "The two strings should be the same.")
// }
