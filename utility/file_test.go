package utility

import (
	"testing"
)

func TestReadLines(t *testing.T) {
	lines := ReadLines("./testdata/1")

	if len(lines) != 3 {
		t.Errorf("wrong lines: except %d, actual %d", 7, len(lines))
	}

}

func TestReadLinesToBytes(t *testing.T) {
	bytes := ReadLinesToBytes("./testdata/1")

	if len(bytes) != 24 {
		t.Errorf("ReadLinesToBytes failed: except %d, actual %d", 24, len(bytes))
	}
}
