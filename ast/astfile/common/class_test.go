package common

import "testing"

func TestFindClassOrInterfaceNames(t *testing.T) {
	ast.generateAstdataFromAstFile("../../testdata/ast")

	ast.findClassOrInterfaceNames()
	if ast.GetInterfacesNames()[0] != "X" {
		t.Error("find wrong class or interface name")
	}
}
