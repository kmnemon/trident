package common

import (
	"strings"
	"trident/ast/object"
)

func (ast *Ast) findMethodNames(p *object.Project) {
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

		ast.findMethodNamesInClass(data, classOrInterfaceName, p)

		return false
	}

	ast.transverseAny(ast.astData, filter, operate)
}

func (ast *Ast) findMethodNamesInClass(data any, className string, p *object.Project) {
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
		generateMethods(ast.packageName, className, methodName, p)
		return false
	}

	ast.transverseAny(data, filter, operate)
}

func generateMethods(packname, cname, mname string, p *object.Project) {
	m := object.NewMethod(mname)
	p.Packages[packname].Classes[cname].Methods[m.Name] = *m
}
