package helper

import (
	"net"
	"os"
	"strings"
)

// GetIPAddress gets the ip address of the host
func GetIPAddress() (addr string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "<???>"
	}

	for _, a := range addrs {
		if ip, ok := a.(*net.IPNet); ok && !ip.IP.IsLoopback() { // skip loopback
			if ip.IP.To4() != nil {
				return ip.IP.String()
			}
		}
	}

	return "<???>"
}

// GetDataCenterName get data center name based on DATACENTER environment variable
func GetDataCenterName() Datacenter {
	datacenter := DC_GCP
	dc := os.Getenv("DATACENTER")

	if strings.HasPrefix(dc, "gcp") {
		datacenter = DC_GCP
	}
	return datacenter
}
