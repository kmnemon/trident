package common

import (
	"testing"
)

func TestFindClassOrInterfaceNames(t *testing.T) {
	a.generateAstdataFromAstFile("../../testdata/ast")

	a.findPackageName(p)
	a.findClassOrInterfaceNames(p)

	if p.Packages[a.packageName].Classes["X"].Name != "X" {
		t.Error("find wrong classes")
	}

	if p.Packages[a.packageName].Classes["X"].VariableCount != 1 {
		t.Error("find wrong number of classes variable")
	}

}
