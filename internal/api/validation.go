package api

import (
	"net"
)

func IpPatternValidation(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	} else {
		return true
	}
}
