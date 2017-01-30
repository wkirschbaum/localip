package localip

import (
	"net"
)

func getLocalIP(fallback string) string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return fallback
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return fallback
}
