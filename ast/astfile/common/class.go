package common

import (
	"strings"
	"trident/ast/object"
)

func (ast *Ast) findClassOrInterfaceNames(p *object.Project) {
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
			generateClasses(ast.packageName, name, true, p)
		} else {
			generateClasses(ast.packageName, name, false, p)
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

func generateClasses(packname string, cname string, isInterface bool, p *object.Project) {
	c := object.NewClass(cname, isInterface)
	p.Packages[packname].Classes[c.Name] = *c
}
