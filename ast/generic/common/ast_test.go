package common

import (
	"fmt"
	"os"
	"testing"
)

var ast Ast[any]

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
	var ast Ast[any]
	ast.init()
	ast.generateAstdataFromAstFile("../testdata/ast2")
	fmt.Println("hello")
}

func TestTransverseMap(t *testing.T) {
	ast.transverseMap(ast.astdata, nil, nil)

}
