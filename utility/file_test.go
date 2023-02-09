package utility

import (
	"fmt"
	"testing"
)

func TestReadLines(t *testing.T) {
	lines := ReadLines("../testdata/1")

	if len(lines) != 3 {
		t.Errorf("wrong lines: except %d, actual %d", 3, len(lines))
	}

}

func TestTest(t *testing.T) {
	s := []string{"abc", "bcd", "efg"}
	b := s
	s = s[:len(s)-1]

	s[0] = "123"
	b[1] = "456"
	b[2] = "789"
	fmt.Println(s)
	fmt.Println(b)

}
