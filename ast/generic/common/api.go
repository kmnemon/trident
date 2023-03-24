package common

type AST interface {
	// GetAllFunctionsName() []string
	// GetAllClassesName() []string
}

func NewAstData(path string) AST {
	var ast Ast[any]
	ast.init()
	ast.generateAstdataFromAstFile(path)
	return &ast
}
