package protocol

import "net"

// ServerRequest is a struct to reference a Request and it's origin.
type ServerRequest struct {
	Request    *Request
	ClientAddr *net.UDPAddr
}

// Request is the base struct for all requests sent to the server
type Request struct {
	ProtocolVersion Version
	RequestType     RequestType
	RequestData     RequestData
	Decoded         bool
}

// RequestData byte array
type RequestData [1024]byte

// RequestType enum
type RequestType uint16

// Version enum for Protocol Version
type Version uint16

// Message is the base struct for messages
type Message struct {
	Length       uint64
	DecodedValue string
}
