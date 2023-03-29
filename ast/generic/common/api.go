package common

func NewAstData(path string) *Ast {
	var ast Ast
	ast.init()
	ast.generateAstdataFromAstFile(path)
	return &ast
}

func (ast *Ast) GetPackageName() string {
	return ast.findPackageName()
}

func (ast *Ast) GetClassOrInterfaceNames() []string {
	ast.findClassOrInterfaceNames()
	return ast.classNames
}

func (ast *Ast) GetMethodNames() []string {
	ast.findMethodNames()
	return ast.methodNames
}
