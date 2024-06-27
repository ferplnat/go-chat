package gui

import (
	"strings"
	"testing"
)

func TestCreateFrame(t *testing.T) {
	var mockFrame string
	topAndBottom := strings.Repeat(string(rune(0x2588)), 10) + "\n"
	middle := string(rune(0x2588)) + strings.Repeat(" ", 8) + string(rune(0x2588)) + "\n"

	mockFrame += topAndBottom
	mockFrame += strings.Repeat(middle, 8)
	mockFrame += topAndBottom

	borderChars := BorderChars{
		Top:    rune(0x2588),
		Bottom: rune(0x2588),
		Left:   rune(0x2588),
		Right:  rune(0x2588),

		CornerTopLeft:     rune(0x2588),
		CornerTopRight:    rune(0x2588),
		CornerBottomLeft:  rune(0x2588),
		CornerBottomRight: rune(0x2588),
	}

	frameOptions := FrameOptions{
		BorderChars: &borderChars,
		Width:       10,
		Height:      10,
	}

	testFrame := CreateFrame(frameOptions)

	if mockFrame != testFrame {
		t.Log("Frames are not equal\n" +
			"Mock:\n" +
			mockFrame + "\n" +
			"Generated:\n" +
			testFrame)

		t.Fail()
	}
}
