package terminalesc

import "testing"

func TestStringToCodeSequence(t *testing.T) {
	testString := "testingthisstringto do things like "
	codeSequence := StringToCodeSequence(testString).String()

	if testString != codeSequence {
		t.Fatalf("Strings do not match. Bytes: \nTest: %b\nConverted: %b", []byte(testString), []byte(codeSequence))
	}
}
