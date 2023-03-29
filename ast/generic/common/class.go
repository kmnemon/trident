package common

import "strings"

func (ast *Ast) findClassOrInterfaceNames() {
	filter := func(s string) bool {
		if strings.Contains(s, "Type=ClassOrInterfaceDeclaration") {
			return true
		} else {
			return false
		}
	}

	operate := func(data any) bool {
		name, err := findNameOperate(data, "SimpleName")
		if err != nil {
			panic("can not find classes or interface name in class declaration")
		}
		if isInterface(data) {
			ast.interfaceNames = append(ast.interfaceNames, name)
		} else {
			ast.classNames = append(ast.classNames, name)
		}
		return false
	}

	ast.transverseAny(ast.astData, filter, operate)
}

func isInterface(data any) bool {
	if data.(map[string]any)["isInterface"].(string) == "true" {
		return true
	} else {
		return false
	}
}
