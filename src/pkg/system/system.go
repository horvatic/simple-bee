package system

import (
	"net"
	"os"
	"strings"
)

func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return hostname
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	panic("Can't find IP")
}

func GetTags() []string {
	tags := os.Getenv("SIMPLE_BEE_TAGS")
	if tags == "" {
		return nil
	}
	return strings.Split(tags, ",")
}
