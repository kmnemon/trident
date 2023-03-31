package common

import (
	"testing"
)

func TestFindClassOrInterfaceNames(t *testing.T) {
	ast.generateAstdataFromAstFile("../../testdata/ast")

	ast.findPackageName(p)
	ast.findClassOrInterfaceNames(p)

	if p.Packages[ast.packageName].Classes["X"].Name != "X" {
		t.Error("find wrong classes")
	}

}
