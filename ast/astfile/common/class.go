package common

import (
	"log"
	"strings"
	"trident/ast/object"
)

func (a *ast) findClassOrInterfaceNames(p *object.Project) {
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
			log.Fatal("can not find classes or interface name in class declaration")
		}
		if isInterface(data) {
			generateClasses(a.packageName, name, true, p)
		} else {
			generateClasses(a.packageName, name, false, p)
		}

		a.findVariables(data, name, p)
		return false
	}

	a.transverseAny(a.astData, filter, operate)
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
	p.Packages[packname].Classes[c.Name] = c

}

func (a *ast) findVariables(data any, cname string, p *object.Project) {
	filter := func(s string) bool {
		if strings.Contains(s, "variable(Type=VariableDeclarator)") {
			return true
		} else {
			return false
		}
	}

	operate := func(data any) bool {
		p.Packages[a.packageName].Classes[cname].VariableCount += 1

		return false
	}

	a.transverseAny(data, filter, operate)
}
