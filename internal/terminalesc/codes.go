// Package terminalesc contains terminal escape codes and helpers for working with them
package terminalesc

import "fmt"

// Code is a rune type that represents a terminal escape code
type Code rune

// String returns the string representation of the Code
func (c Code) String() string {
	return string(c)
}

// Code returns the Code as a rune
func (c Code) Code() rune {
	return rune(c)
}

// BEL is a rune that represents a Terminal Bell
var BEL = Code(0x07)

// BS is a rune that represents a Backspace
var BS = Code(0x08)

// HT is a rune that represents a Horizontal Tab
var HT = Code(0x09)

// LF is a rune that represents a Line Feed
var LF = Code(0x0A)

// VT is a rune that represents a Vertical Tab
var VT = Code(0x0B)

// FF is a rune that represents a Form Feed
var FF = Code(0x0C)

// CR is a rune that represents a Carriage Return
var CR = Code(0x0D)

// ESC is a rune that represents an Escape character
var ESC = Code(0x1B)

// DEL is a rune that represents a Delete character
var DEL = Code(0x7F)

// CodeSequence is an array type that represents a Terminal Escape sequence.
type CodeSequence []Code

func (c CodeSequence) String() string {
	return string(c)
}

// CreateSequence creates a CodeSequence from an array of Code elements.
func CreateSequence(code []Code) CodeSequence {
	return CodeSequence(code)
}

// CreateSequenceWithEsc creates a CodeSequence from an array of Code elements and prepends the ESC Code.
func CreateSequenceWithEsc(code []Code) CodeSequence {
	newCode := make([]Code, len(code)+1)
	newCode[0] = ESC
	copy(newCode[1:], code)

	return CodeSequence(newCode)
}

// CursorMoveHome is a Sequence that moves the cursor to 0,0 when printed
var CursorMoveHome = CreateSequenceWithEsc(StringToCodeSequence("[H"))

// CursorMove is a Sequence that moves the cursor to the given line and column int.
func CursorMove(line int, column int) CodeSequence {
	seq := StringToCodeSequence(fmt.Sprintf("[%d;%dH", line, column))
	return CreateSequenceWithEsc(seq)
}

// CursorMoveUp is a Sequence that moves the cursor up by int count.
func CursorMoveUp(count int) CodeSequence {
	seq := StringToCodeSequence(fmt.Sprintf("[%dA", count))
	return CreateSequenceWithEsc(seq)
}

// CursorMoveDown is a Sequence that moves the cursor down by int count.
func CursorMoveDown(count int) CodeSequence {
	seq := StringToCodeSequence(fmt.Sprintf("[%dB", count))
	return CreateSequenceWithEsc(seq)
}

// CursorMoveRight is a Sequence that moves the cursor right by int count.
func CursorMoveRight(count int) CodeSequence {
	seq := StringToCodeSequence(fmt.Sprintf("[%dC", count))
	return CreateSequenceWithEsc(seq)
}

// CursorMoveLeft is a Sequence that moves the cursor left by int count.
func CursorMoveLeft(count int) CodeSequence {
	seq := StringToCodeSequence(fmt.Sprintf("[%dD", count))
	return CreateSequenceWithEsc(seq)
}

// CursorMoveToColumn moves the cursor to the given column int.
func CursorMoveToColumn(column int) CodeSequence {
	seq := StringToCodeSequence(fmt.Sprintf("[%dG", column))
	return CreateSequenceWithEsc(seq)
}

// EraseInDisplay terminal escape sequence.
var EraseInDisplay = CreateSequenceWithEsc(StringToCodeSequence("[J"))

// EraseFromCursorToScreenEnd terminal escape sequence.
var EraseFromCursorToScreenEnd = CreateSequenceWithEsc(StringToCodeSequence("[0J"))

// EraseFromCursorToScreenBeginning terminal escape sequence.
var EraseFromCursorToScreenBeginning = CreateSequenceWithEsc(StringToCodeSequence("[1J"))

// EraseScreen terminal escape sequence.
var EraseScreen = CreateSequenceWithEsc(StringToCodeSequence("[2J"))

// EraseSavedLines terminal escape sequence.
var EraseSavedLines = CreateSequenceWithEsc(StringToCodeSequence("[3J"))

// EraseInLine terminal escape sequence.
var EraseInLine = CreateSequenceWithEsc(StringToCodeSequence("[K"))

// EraseFromCursorToLineEnd terminal escape sequence.
var EraseFromCursorToLineEnd = CreateSequenceWithEsc(StringToCodeSequence("[0K"))

// EraseFromCursorToLineBeginning terminal escape sequence.
var EraseFromCursorToLineBeginning = CreateSequenceWithEsc(StringToCodeSequence("[1K"))

// EraseLine terminal escape sequence.
var EraseLine = CreateSequenceWithEsc(StringToCodeSequence("[2K"))

// SetTextFormat returns an escape sequence to format superseding text. An error is returned if the format is invalid.
func SetTextFormat(format string) (CodeSequence, error) {
	formatStr := "[%dm"
	var str string

	switch format {
	case "Bold", "bold":
		str = fmt.Sprintf(formatStr, 1)

	case "Dim", "dim", "Faint", "faint":
		str = fmt.Sprintf(formatStr, 2)

	case "Italic", "italic":
		str = fmt.Sprintf(formatStr, 3)

	case "Underline", "underline":
		str = fmt.Sprintf(formatStr, 4)

	case "Blinking", "blinking":
		str = fmt.Sprintf(formatStr, 5)

	case "Inverse", "inverse", "Reverse", "reverse":
		str = fmt.Sprintf(formatStr, 7)

	case "Hidden", "hidden", "Invisible", "invisible":
		str = fmt.Sprintf(formatStr, 8)

	case "Strikethrough", "strikethrough":
		str = fmt.Sprintf(formatStr, 9)

	default:
		return nil, &FormatError{}
	}

	return CreateSequenceWithEsc(StringToCodeSequence(str)), nil
}

// ResetTextFormat returns an escape sequence that denotes the end of a given format specifier.
func ResetTextFormat(format string) (CodeSequence, error) {
	formatStr := "[%dm"
	var str string

	switch format {
	case "Bold", "bold",
		"Dim", "dim", "Faint", "faint":
		str = fmt.Sprintf(formatStr, 22)

	case "Italic", "italic":
		str = fmt.Sprintf(formatStr, 23)

	case "Underline", "underline":
		str = fmt.Sprintf(formatStr, 24)

	case "Blinking", "blinking":
		str = fmt.Sprintf(formatStr, 25)

	case "Inverse", "inverse", "Reverse", "reverse":
		str = fmt.Sprintf(formatStr, 27)

	case "Hidden", "hidden", "Invisible", "invisible":
		str = fmt.Sprintf(formatStr, 28)

	case "Strikethrough", "strikethrough":
		str = fmt.Sprintf(formatStr, 29)

	default:
		return nil, &FormatError{}
	}

	return CreateSequenceWithEsc(StringToCodeSequence(str)), nil
}
