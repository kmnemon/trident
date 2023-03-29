package common

import (
	"os"
	"testing"
	"trident/utility"
)

var ast Ast

func TestMain(m *testing.M) {
	setup()
	extcode := m.Run()
	teardown()

	os.Exit(extcode)
}

func setup() {
	ast.init()
	ast.generateAstdataFromAstFile("../testdata/ast2")
}

func teardown() {
}

func TestGenerateAstdataFromAstFile(t *testing.T) {
	ast.generateAstdataFromAstFile("../testdata/ast2")

	var r []string = make([]string, 0)
	for k, v := range ast.astData {
		r = append(r, k)
		for k1 := range v.(map[string]any) {
			r = append(r, k1)
		}
	}

	expect := []string{"root(Type=CompilationUnit)", "import", "good", "types"}

	if !utility.EqualSliceHelper(expect, r) {
		t.Error("generate astdata wrong")
	}
}

func TestTransverseAny(t *testing.T) {
	var r string
	filter := func(s string) bool {
		if s == "good" {
			return true
		} else {
			return false
		}
	}

	operate := func(a any) bool {
		r = a.(string)
		return false
	}

	ast.transverseAny(ast.astData, filter, operate)

	if r != "what" {
		t.Error("transverse any wrong")
	}

}

func TestTransverseAnyContainSlice(t *testing.T) {
	var r []any
	filter := func(s string) bool {
		if s == "types" {
			return true
		} else {
			return false
		}
	}

	operate := func(a any) bool {
		r = a.([]any)
		return false
	}

	ast.transverseAny(ast.astData, filter, operate)

	expect := []any{"type(Type=ClassOrInterfaceDeclaration)", "type2"}
	if !utility.EqualSliceHelper(expect, r) {
		t.Error("transverse any contain slice wrong")
	}

}

func TestFindPackageName(t *testing.T) {
	ast.generateAstdataFromAstFile("../testdata/ast")

	name := ast.findPackageName()
	if name != "ast" {
		t.Error("find wrong package name")
	}

}
