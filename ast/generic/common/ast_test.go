package common

import "testing"

func TestGenerateAstdataFromAstFile(t *testing.T) {
	var ast Ast
	ast.init()
	ast.generateAstdataFromAstFile("../testdata/ast2")
}
