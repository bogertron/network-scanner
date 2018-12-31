package main

import (
	"log"
	"net"
	"time"

	"github.com/mdlayher/arp"
)

//
// successful resonse for an ip address request
//
type ARPResponse struct {
	hwAddress net.HardwareAddr
}

func CheckARPRequest(ipAddr string) (*ARPResponse, error) {
	ifaces, err := net.Interfaces()
	ip, err := net.LookupIP(ipAddr)

	for _, i := range ifaces {
		if i.Name != "lo" {

			if err != nil {
				log.Panic("Unable to lookup ip [", ipAddr, "]")
			} else {
				arpHW, err := arp.Dial(&i)

				if err != nil {
					log.Panic("Failed to resolve hw for [", ipAddr, "]")
				} else {
					defer arpHW.Close()
					for _, ipAddrIt := range ip {
						t := time.Now()
						t = t.Add(time.Second * 30)
						response := new(ARPResponse)
						err := arpHW.SetDeadline(t)
						if err == nil {
							arpRes, err := arpHW.Resolve(ipAddrIt)
							if arpRes != nil && len(arpRes) > 0 {
								log.Print("ARPRES ", arpRes)
							}
							// Resolve can fail to resolve the address, but return an empty value
							// instead of setting the err value
							if err != nil && len(response.hwAddress) > 0 {
								log.Print("ARP resolve failed: ", err)
								return response, nil
							}
						} else {
							log.Print("Failed to set ARP deadline")
						}
					}
				}
			}
		}
	}
	return nil, err
}
