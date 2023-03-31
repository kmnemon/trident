package common

import (
	"os"
	"testing"
	"trident/ast/object"
	"trident/utility"
)

var a ast
var p *object.Project

func TestMain(m *testing.M) {
	setup()
	extcode := m.Run()
	teardown()

	os.Exit(extcode)
}

func setup() {
	a.init()
	a.generateAstdataFromAstFile("../../testdata/ast2")
	p = object.NewProject("1")
}

func teardown() {
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

	a.transverseAny(a.astData, filter, operate)

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

	a.transverseAny(a.astData, filter, operate)

	expect := []any{"type(Type=ClassOrInterfaceDeclaration)", "type2"}
	if !utility.EqualSliceHelper(expect, r) {
		t.Error("transverse any contain slice wrong")
	}

}
