package common

import (
	"testing"
)

func TestAstFileToObject(t *testing.T) {
	newAstFileAddToObject("../../testdata/ast", p)

	if p.Packages["ast"].Name != "ast" {
		t.Error("project get package name wrong  ")
	}

	if len(p.Packages["ast"].Classes) != 1 {
		t.Error("project get class  wrong  ")
	}

}

func TestGenerateObjects(t *testing.T) {
	p := GenerateObjects("../../testdata/")

	if len(p.Packages) != 1 {
		t.Error("project get wrong package number")
	}
}

func TestFindPackageName(t *testing.T) {
	a.generateAstdataFromAstFile("../../testdata/ast")

	name := a.findPackageName(p)
	if name != "ast" {
		t.Error("find wrong package name")
	}

}
