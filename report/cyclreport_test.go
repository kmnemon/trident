package report

import (
	"reflect"
	"testing"
	"trident/utility"
)

func TestFileToMap(t *testing.T) {
	readOriginalReport()
}

func TestStringToMapUsingCycls(t *testing.T) {
	a := []string{
		`D:\10_Code\DTO.java:184: CyclomaticComplexity: The method 'sdf' has a cyclomatic complexity of 2.`,
		`D:\10_Code\DTO.java:185: CyclomaticComplexity: The method 'sdf' has a cyclomatic complexity of 2.`,
		`D:\10_Code\DTO3.java:186: CyclomaticComplexity: The method 'sdf' has a cyclomatic complexity of 5.`,
	}

	b := stringToMapUsingCycls(a)

	expect := map[int][]string{
		2: {`D:\10_Code\DTO.java:184: CyclomaticComplexity: The method 'sdf' has a cyclomatic complexity of 2.`,
			`D:\10_Code\DTO.java:185: CyclomaticComplexity: The method 'sdf' has a cyclomatic complexity of 2.`},
		5: {`D:\10_Code\DTO3.java:186: CyclomaticComplexity: The method 'sdf' has a cyclomatic complexity of 5.`},
	}

	if !reflect.DeepEqual(expect, b) {
		t.Error("string to map using cycls have something wrong")
	}
}

func TestFilterContentOut(t *testing.T) {
	a := []string{
		"abc", "bcd", "def",
	}

	a = filterContentOut(a, "bcd")

	expect := []string{
		"abc", "def",
	}

	if !utility.EqualSliceHelper(expect, a) {
		t.Error("filter content out failed")
	}

}
