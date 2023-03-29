package object

import "trident/ast/generic/common"

type Project struct {
	packages map[string]Package
}

func GenerateObjects(astfilepath string) *Project {
	var p Project
	p.packages = make(map[string]Package)
	return &p
}

func (p *Project) generateObject(path string) {
	ast := common.NewAstData(path)
	packageName := ast.GetPackageName()
	pack := NewPackage(packageName)

	p.packages[packageName] = *pack

}
