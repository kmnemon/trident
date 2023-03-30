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
		classOrInterfaceName, err := findNameOperate(data, "SimpleName")
		if err != nil {
			panic("can not find classes or interface name in class declaration")
		}

		if isInterface(data) {
			classOrInterfaceName = "i." + classOrInterfaceName
		} else {
			classOrInterfaceName = "c." + classOrInterfaceName
		}

		ast.findMethodNamesInClass(data, classOrInterfaceName)

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
		ast.methodsNames = append(ast.methodsNames, className+"."+methodName)
		return false
	}

	ast.transverseAny(data, filter, operate)
}
