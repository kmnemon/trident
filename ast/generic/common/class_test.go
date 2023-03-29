package common

import "testing"

func TestFindClassOrInterfaceNames(t *testing.T) {
	ast.generateAstdataFromAstFile("../testdata/ast")

	ast.findClassOrInterfaceNames()
	if ast.packageName != "ast" {
		t.Error("find wrong class or interface name")
	}
}
