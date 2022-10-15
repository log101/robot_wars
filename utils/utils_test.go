package utils

import "testing"

func Test_Random3(t *testing.T) {
	got := Random3()

	if got > 3 || got < 0 {
		t.Error("random3 should be between 0-3")
	}
}
