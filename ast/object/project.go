package object

import (
	"fmt"
	"strings"
)

type Project struct {
	Name     string
	Packages map[string]*Package
}

func NewProject(name string) *Project {
	var p Project
	p.Name = name
	p.Packages = make(map[string]*Package)
	return &p
}

// i.interfacename.methodname  c.classname.methodname
func splitClassOrInterfaceName(name string) (string, string, string) {
	s := strings.Split(name, ".")
	return s[0], s[1], s[2]
}

func (p *Project) ToString() string {
	var str string
	str += fmt.Sprintf("Project: %s\n", p.Name)
	for _, v := range p.Packages {
		str += fmt.Sprintf("    Package: %s\n", v.Name)
		str += fmt.Sprintln("        Class or Interface:")
		for _, v1 := range v.Classes {
			var ic string
			if v1.IsInterface {
				ic = "Interface"
			} else {
				ic = "Class"
			}
			str += fmt.Sprintln("          " + v1.Name + " : " + ic)
			str += fmt.Sprintln("              Method:")
			for _, v2 := range v1.Methods {
				str += fmt.Sprintln("                 " + v2.Name)
			}
		}
	}

	return str
}
