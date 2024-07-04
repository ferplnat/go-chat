package protocol

import (
	"testing"
)

var MockMessageString = "test message"
var MockMessage = CreateMessage(MockMessageString)
var MockMessageRequest = CreateMessageRequest(MockMessage)

func TestDecode(t *testing.T) {
	t.Log("Decoding message")
	MockMessageRequest.Decode()
}

func TestCreateMessage(t *testing.T) {
	message := CreateMessage(MockMessageString)

	t.Logf("Reported length: %d", message.Length)
	t.Logf("Actual length: %d", len(MockMessageString))
}

func TestGetMessage(t *testing.T) {
	MockMessageRequest.Decode()
	message := MockMessageRequest.GetMessage()
	t.Logf("Got message: \"%s\"", message.DecodedValue)

	if message.DecodedValue != MockMessageString {
		t.Logf("DecodedValue bytes: %b", []byte(message.DecodedValue))
		t.Logf("MockMessageString bytes: %b", []byte(MockMessageString))
		t.Fail()
	}
}

func TestGetSetProtocolVersion(t *testing.T) {
	testRequest := Request{}
	testVersion := Version(2)
	t.Logf("Protocol version: (%d)", testVersion)

	testRequest.SetProtocolVersion(testVersion)
	if testRequest.GetProtocolVersion() != testVersion {
		t.Fatalf("Tested version (%d) does not match resolved version (%d)", testVersion, testRequest.GetProtocolVersion())
	}

}

func TestGetSetRequestType(t *testing.T) {
	testRequest := Request{}

	testType := RequestType(5)
	t.Logf("RequestType: (%d)", testType)

	testRequest.SetRequestType(testType)
	if testRequest.GetRequestType() != testType {
		t.Fatalf("Tested RequestType (%d) does not match resolved RequestType (%d)", testType, testRequest.GetRequestType())
	}
}
