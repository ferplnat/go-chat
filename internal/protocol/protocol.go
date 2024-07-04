// Package protocol contains the protocol definition for the communication between the client and the server.
package protocol

import "encoding/binary"

// Decode unpacks the byte data into the Request struct
func (r *Request) Decode() {
	r.ProtocolVersion = r.GetProtocolVersion()
	r.RequestType = r.GetRequestType()

	r.Decoded = true
}

// GetProtocolVersion returns the Version in the raw RequestData
func (r *Request) GetProtocolVersion() Version {
	return Version(binary.BigEndian.Uint16(r.RequestData[0:2]))
}

// SetProtocolVersion sets the Version in the raw RequestData
func (r *Request) SetProtocolVersion(version Version) {
	binary.BigEndian.PutUint16(r.RequestData[0:2], uint16(version))
}

// GetRequestType returns the RequestType in the raw RequestData
func (r *Request) GetRequestType() RequestType {
	return RequestType(binary.BigEndian.Uint16(r.RequestData[2:4]))
}

// SetRequestType sets the RequestType in the raw RequestData
func (r *Request) SetRequestType(requestType RequestType) {
	binary.BigEndian.PutUint16(r.RequestData[2:4], uint16(requestType))
}

// GetMessage will decode if not already decoded and extract the message into the Message struct
func (r *Request) GetMessage() Message {
	var message Message
	if !r.Decoded {
		r.Decode()
	}

	message.Length = binary.BigEndian.Uint64(r.RequestData[4:12])
	message.DecodedValue = string(r.RequestData[12 : 13+message.Length-1])

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

	binary.BigEndian.PutUint64(request.RequestData[4:12], message.Length)

	for i, v := range message.DecodedValue {
		request.RequestData[12+i] = byte(v)
	}

	return request
}
