package gui

import (
	"reflect"
	"strings"
)

type BorderChars struct {
	Top    rune
	Bottom rune
	Left   rune
	Right  rune

	CornerTopLeft     rune
	CornerTopRight    rune
	CornerBottomLeft  rune
	CornerBottomRight rune
}

var defaultBorderChars = BorderChars{
	Top:    rune(0x2588),
	Bottom: rune(0x2588),
	Left:   rune(0x2588),
	Right:  rune(0x2588),

	CornerTopLeft:     rune(0x2588),
	CornerTopRight:    rune(0x2588),
	CornerBottomLeft:  rune(0x2588),
	CornerBottomRight: rune(0x2588),
}

func (b *BorderChars) SetDefault() {
	*b = defaultBorderChars
}

type FrameOptions struct {
	BorderChars *BorderChars
	Width       int
	Height      int
}

var defaultFrameOptions = FrameOptions{
	BorderChars: &defaultBorderChars,
	Width:       75,
	Height:      50,
}

func (f *FrameOptions) SetDefault() {
	*f = defaultFrameOptions
}

func isValueSet(x interface{}) bool {
	return reflect.New(reflect.TypeOf(x)) != reflect.ValueOf(x)
}

func (b *BorderChars) mergeDefaults() {
	if isValueSet(b.Top) {
		b.Top = defaultBorderChars.Top
	}
}

func (f *FrameOptions) mergeDefaults() {
}

func CreateFrame(f FrameOptions) string {
	totalSize := f.Height * (f.Width + 1) // Add room for line breaks
	frame := make([]rune, 0, totalSize)

	top := string(f.BorderChars.CornerTopLeft) +
		strings.Repeat(string(f.BorderChars.Top), f.Width-2) +
		string(f.BorderChars.CornerTopRight) + "\n"

	middle := string(f.BorderChars.Left) +
		strings.Repeat(string(" "), f.Width-2) +
		string(f.BorderChars.Right) + "\n"
	middle = strings.Repeat(middle, f.Height-2)

	bottom := string(f.BorderChars.CornerBottomLeft) +
		strings.Repeat(string(f.BorderChars.Bottom), f.Width-2) +
		string(f.BorderChars.CornerBottomRight) + "\n"

	frame = append(frame, []rune(top)...)
	frame = append(frame, []rune(middle)...)
	frame = append(frame, []rune(bottom)...)

	return string(frame)
}
