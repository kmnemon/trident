package utility

import (
	"testing"
)

func TestReadLines(t *testing.T) {
	lines := ReadLines("../reportfile/1")

	if len(lines) != 7 {
		t.Errorf("wrong lines: except %d, actual %d", 7, len(lines))
	}

}
