// client
package main

import (
	"github.com/ferplnat/go-chat/internal/protocol"
	"net"
)

func main() {
	addr := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 12345,
		Zone: "",
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return
	}

	message := []byte("hello")
	conn.Write(message)
}
