package object

type Package struct {
	Name    string
	Classes map[string]Class
}

func NewPackage(name string) *Package {
	var pack Package
	pack.Name = name
	pack.Classes = make(map[string]Class)
	return &pack
}
