package common

import "strings"

func (ast *Ast) findMethodNames() {
	filter := func(s string) bool {
		if strings.Contains(s, "Type=ClassOrInterfaceDeclaration") {
			return true
		} else {
			return false
		}
	}

	operate := func(data any) bool {
		className, err := findNameOperate(data, "SimpleName")
		if err != nil {
			panic("can not find classes or interface name in class declaration")
		}

		ast.findMethodNamesInClass(data, className)

		return false
	}

	ast.transverseAny(ast.astData, filter, operate)
}

func (ast *Ast) findMethodNamesInClass(data any, className string) {
	filter := func(s string) bool {
		if strings.Contains(s, "Type=MethodDeclaration") {
			return true
		} else {
			return false
		}
	}

	operate := func(data any) bool {
		methodName, err := findNameOperate(data, "SimpleName")
		if err != nil {
			panic("can not find classes or interface name in class declaration")
		}
		ast.methodNames = append(ast.methodNames, className+"."+methodName)
		return false
	}

	ast.transverseAny(data, filter, operate)
}
