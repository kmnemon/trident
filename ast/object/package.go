package object

type Package struct {
	name       string
	classes    map[string]Class
	interfaces map[string]Interface
}

func NewPackage(name string) *Package {
	var p Package
	p.name = name
	p.classes = make(map[string]Class)
	p.interfaces = make(map[string]Interface)
	return &p
}
