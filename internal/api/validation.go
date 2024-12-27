package api

import (
	"net"
)

func Ip_pattern_validation(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	} else {
		return true
	}

}
