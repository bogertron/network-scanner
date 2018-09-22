package main

import (
	"log"

	"golang.org/x/net/icmp"
)

// Request to see if the address passed in can resolve
type ICMPRequest struct {
	ipAddress string
}

type ICMPResponse struct {
	ipAddress string
	success   bool
}

func CheckRequest(request *ICMPRequest) *ICMPResponse {
	conn, err := icmp.ListenPacket("ip4:icmp", request.ipAddress)
	var response = new(ICMPResponse)
	response.ipAddress = request.ipAddress
	if err != nil {
		response.success = false
		log.Print("Connect failed ", err)
	} else {
		response.success = true
	}

	if conn != nil {
		log.Printf("Connected to ", request.ipAddress)
		conn.Close()
	}

	return response
}
