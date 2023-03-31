package common

import "testing"

func TestFindMethodNames(t *testing.T) {
	ast.generateAstdataFromAstFile("../../testdata/ast")

	ast.findPackageName(p)
	ast.findClassOrInterfaceNames(p)
	ast.findMethodNames(p)

	if len(p.Packages[ast.packageName].Classes["X"].Methods) != 2 {
		t.Error("find wrong methods")
	}

}
