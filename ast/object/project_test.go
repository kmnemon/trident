package object

import (
	"testing"
)

func TestAstFileToObject(t *testing.T) {
	p := NewProject("1")
	p.AstFileToObject("../testdata/ast")

	if p.GetPackages()["ast"].name != "ast" {
		t.Error("project get package name wrong  ")
	}

	if len(p.GetPackages()["ast"].classes) != 0 {
		t.Error("project get class  wrong  ")
	}

	if len(p.GetPackages()["ast"].interfaces) != 1 {
		t.Error("project get interface  wrong  ")
	}

	if len(p.GetPackages()["ast"].interfaces["X"].methods) != 2 {
		t.Error("project get method  wrong  ")
	}

}

func TestGenerateObjects(t *testing.T) {
	p := GenerateObjects("../testdata/")

	if len(p.GetPackages()) != 2 {
		t.Error("project get wrong package number")
	}
}
