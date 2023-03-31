package common

import "testing"

func TestFindMethodNames(t *testing.T) {
	a.generateAstdataFromAstFile("../../testdata/ast")

	a.findPackageName(p)
	a.findClassOrInterfaceNames(p)
	a.findMethodNames(p)

	if len(p.Packages[a.packageName].Classes["X"].Methods) != 2 {
		t.Error("find wrong methods")
	}

	if p.Packages[a.packageName].Classes["X"].Methods["fn2"].ParameterCount != 2 {
		t.Error("find wrong method parameter")
	}

}
