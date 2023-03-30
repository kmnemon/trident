package common

func NewAstData(path string) *Ast {
	var ast Ast
	ast.init()
	ast.generateAstdataFromAstFile(path)

	ast.findPackageName()
	ast.findClassOrInterfaceNames()
	ast.findMethodNames()
	return &ast
}
