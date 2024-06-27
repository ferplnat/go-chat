package terminalesc

import (
	"fmt"
	"testing"
)

func TestMoveHome(t *testing.T) {
	fmt.Print(CursorMoveHome)
}
