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
		buf := make([]byte, 1024)
		_, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			return
		}

		fmt.Printf("hello %s", remoteAddr.String())
		clientConns.Store(remoteAddr.String(), &remoteAddr)
	}
}
