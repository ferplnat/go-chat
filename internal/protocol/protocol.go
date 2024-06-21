// Package protocol contains the protocol definition for the communication between the client and the server.
package protocol

import "encoding/binary"

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
type RequestType uint64

// Version enum for Protocol Version
type Version uint64

// Message is the base struct for messages
type Message struct {
	Length       uint64
	DecodedValue string
}

// Decode unpacks the byte data into the Request struct
func (r *Request) Decode() {
	r.ProtocolVersion = Version(binary.BigEndian.Uint64(r.RequestData[0:8]))
	r.RequestType = RequestType(binary.BigEndian.Uint64(r.RequestData[8:16]))

	r.Decoded = true
}

// GetMessage will decode if not already decoded and extract the message into the Message struct
func (r *Request) GetMessage() Message {
	var message Message
	if !r.Decoded {
		r.Decode()
	}

	message.Length = binary.BigEndian.Uint64(r.RequestData[16:24])
	message.DecodedValue = string(r.RequestData[24 : 25+message.Length-1])

	return message
}

// CreateMessage creates a Message from a string
func CreateMessage(messageString string) Message {
	return Message{
		Length:       uint64(len(messageString)),
		DecodedValue: messageString,
	}
}

// CreateMessageRequest creates a Request from a Message
func CreateMessageRequest(message Message) Request {
	request := Request{
		ProtocolVersion: 0,
		RequestType:     1,
	}

	binary.BigEndian.PutUint64(request.RequestData[16:24], message.Length)

	for i, v := range message.DecodedValue {
		request.RequestData[24+i] = byte(v)
	}

	return request
}
