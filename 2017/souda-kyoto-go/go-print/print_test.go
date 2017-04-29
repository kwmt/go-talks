package print

import (
	"testing"
)

func TestPrint(t *testing.T) {
	expect := "そうだ Go、京都。"
	actual := Print("そうだ Go、京都。")
	if expect != actual {
		t.Errorf("actual: %s, expect: %s", actual, expect)
	}
}
