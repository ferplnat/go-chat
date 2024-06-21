package protocol

import "testing"

var MockMessageString = "test message"
var MockMessage = CreateMessage(MockMessageString)
var MockMessageRequest = CreateMessageRequest(MockMessage)

func TestDecode(t *testing.T) {
	MockMessageRequest.Decode()
}

func TestCreateMessage(t *testing.T) {
	message := CreateMessage(MockMessageString)

	t.Logf("Reported length: %d", message.Length)
	t.Logf("Acutal length: %d", len(MockMessageString))
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
