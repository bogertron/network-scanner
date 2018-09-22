package main

import (
	"log"
	"net"
	"time"
)

type Host struct {
	Ip            string
	ResponseFound bool
}

// Based on the network string passed in (CIDR), build a list of all potential
// IP addresses on the network
func BuildHosts(cidr string) ([]*Host, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Panic("Unable to parse cidr [", cidr, "]: ", err)
		return nil, err
	}

	var ips []*Host
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); incrementAddress(ip) {
		host := new(Host)
		host.Ip = ip.String()
		host.ResponseFound = false
		ips = append(ips, host)
	}

	return ips, nil
}

//  http://play.golang.org/p/m8TNTtygK0
func incrementAddress(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func main() {
	log.Println("Start of process at ", time.Now().UTC())

	hosts, err := BuildHosts("192.168.1.255/24")

	if err != nil {
		log.Println("Error encountered: ", err)
	}

	log.Println("Found ", len(hosts), " ip addresses")

	for _, host := range hosts {
		log.Println("Address ", host.Ip)
		request := new(ICMPRequest)
		request.ipAddress = host.Ip
		response := CheckRequest(request)
		if response.success {
			log.Print("Received a successful response")
		}
	}
}
