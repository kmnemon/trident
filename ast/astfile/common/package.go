package common

import (
	"log"
	"strings"
	"trident/ast/object"
)

func (a *ast) findPackageName(p *object.Project) string {
	filter := func(s string) bool {
		if strings.Contains(s, "Type=PackageDeclaration") {
			return true
		} else {
			return false
		}
	}

	operate := func(data any) bool {
		name, err := findNameOperate(data, "Name")
		if err != nil {
			log.Fatal("can not find package name in java file")
		}
		a.packageName = name
		generatePackage(name, p)
		return true
	}

	a.transverseAny(a.astData, filter, operate)
	return a.packageName
}

func generatePackage(name string, p *object.Project) {
	_, ok := p.Packages[name]
	if !ok {
		pack := object.NewPackage(name)
		p.Packages[name] = pack
	}
}
