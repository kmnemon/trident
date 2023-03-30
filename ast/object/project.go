package object

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"trident/ast/astfile/common"
)

type Project struct {
	name     string
	packages map[string]Package
}

func NewProject(name string) *Project {
	var p Project
	p.name = name
	p.packages = make(map[string]Package)
	return &p
}

func (p *Project) GetName() string {
	return p.name
}

func (p *Project) GetPackages() map[string]Package {
	return p.packages
}

func GenerateObjects(astfileDir string) *Project {
	p := NewProject(astfileDir)
	files, err := ioutil.ReadDir(astfileDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !strings.HasPrefix(file.Name(), ".") {
			p.AstFileToObject(filepath.Join(astfileDir, file.Name()))
		}

	}

	return p
}

func (p *Project) AstFileToObject(path string) {
	ast := common.NewAstData(path)

	p.generatePackage(ast)
	p.generateClasses(ast)
	p.generateInterfaces(ast)
	p.generateMethods(ast)

}

func (p *Project) generateMethods(ast *common.Ast) {
	for _, v := range ast.GetMethodsNames() {
		isInterface, ciname, mname := splitClassOrInterfaceName(v)
		m := NewMethod(mname)
		if isInterface == "i" {
			p.packages[ast.GetPackageName()].interfaces[ciname].methods[m.GetName()] = *m
		} else {
			p.packages[ast.GetPackageName()].classes[ciname].methods[m.GetName()] = *m
		}
	}
}

func (p *Project) generateInterfaces(ast *common.Ast) {
	for _, v := range ast.GetInterfacesNames() {
		i := NewInterface(v)
		p.packages[ast.GetPackageName()].interfaces[i.GetName()] = *i
	}
}

func (p *Project) generateClasses(ast *common.Ast) {
	for _, v := range ast.GetClassesNames() {
		c := NewClass(v)
		p.packages[ast.GetPackageName()].classes[c.GetName()] = *c
	}
}

func (p *Project) generatePackage(ast *common.Ast) {
	_, ok := p.packages[ast.GetPackageName()]
	if !ok {
		pack := NewPackage(ast.GetPackageName())
		p.packages[ast.GetPackageName()] = *pack
	}

}

// i.interfacename.methodname  c.classname.methodname
func splitClassOrInterfaceName(name string) (string, string, string) {
	s := strings.Split(name, ".")
	return s[0], s[1], s[2]
}

func (p *Project) ToString() string {
	var str string
	str += fmt.Sprintf("Project: %s\n", p.GetName())
	for _, v := range p.packages {
		str += fmt.Sprintf("    Package: %s\n", v.GetName())
		str += fmt.Sprintln("        Class:")
		for _, v1 := range v.GetClasses() {
			str += fmt.Sprintln("          " + v1.GetName())
			str += fmt.Sprintln("              Method:")
			for _, v2 := range v1.GetMethods() {
				str += fmt.Sprintln("                 " + v2.GetName())
			}
		}

		fmt.Sprintln("        Interface:")
		for _, v1 := range v.GetInterfaces() {
			str += fmt.Sprintln("          " + v1.GetName())
			str += fmt.Sprintln("              Method:")
			for _, v2 := range v1.GetMethods() {
				str += fmt.Sprintln("                 " + v2.GetName())
			}
		}
	}

	return str
}
