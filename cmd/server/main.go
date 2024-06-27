// my amateur starting point in Go -- shut up revive
package main

import (
	"fmt"
	"github.com/ferplnat/go-chat/internal/protocol"
	"net"
	"sync"
)

func main() {
	addr := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 12345,
		Zone: "",
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	clientConns := new(sync.Map)

	for {
		request := protocol.Request{}
		serverRequest := protocol.ServerRequest{
			Request: &request,
		}

		_, serverRequest.ClientAddr, err = conn.ReadFromUDP(serverRequest.Request.RequestData[:])
		if err != nil {
			return
		}

		message := serverRequest.Request.GetMessage()

		fmt.Printf("Message from (%s): %s\n", serverRequest.ClientAddr.IP, message.DecodedValue)
		clientConns.Store(serverRequest.ClientAddr.String(), &serverRequest.ClientAddr)
	}
}
