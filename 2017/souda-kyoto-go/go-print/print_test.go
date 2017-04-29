package print

import (
	"testing"
)

func TestPrint(t *testing.T) {
	expect := "そうだ京都Go"
	actual := Print("そうだ京都Go")
	if expect != actual {
		t.Errorf("actual: %s, expect: %s", actual, expect)
	}
}
