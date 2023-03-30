package object

type Package struct {
	name       string
	classes    map[string]Class
	interfaces map[string]Interface
}

func NewPackage(name string) *Package {
	var pack Package
	pack.name = name
	pack.classes = make(map[string]Class)
	pack.interfaces = make(map[string]Interface)
	return &pack
}

func (pack *Package) GetName() string {
	return pack.name
}

func (pack *Package) GetClasses() map[string]Class {
	return pack.classes
}

func (pack *Package) GetInterfaces() map[string]Interface {
	return pack.interfaces
}
