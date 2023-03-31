package common

import (
	"strings"
	"trident/ast/object"
)

func (a *ast) findMethodNames(p *object.Project) {
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

		a.findMethodNamesInClass(data, classOrInterfaceName, p)

		return false
	}

	a.transverseAny(a.astData, filter, operate)
}

func (a *ast) findMethodNamesInClass(data any, className string, p *object.Project) {
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
		generateMethods(a.packageName, className, methodName, p)
		a.findParameters(data, className, methodName, p)
		return false
	}

	a.transverseAny(data, filter, operate)
}

func generateMethods(packname, cname, mname string, p *object.Project) {
	m := object.NewMethod(mname)
	p.Packages[packname].Classes[cname].Methods[m.Name] = m
}

func (a *ast) findParameters(data any, cname, mname string, p *object.Project) {
	filter := func(s string) bool {
		if strings.Contains(s, "parameter(Type=Parameter)") {
			return true
		} else {
			return false
		}
	}

	operate := func(data any) bool {
		p.Packages[a.packageName].Classes[cname].Methods[mname].ParameterCount += 1

		return false
	}

	a.transverseAny(data, filter, operate)
}
