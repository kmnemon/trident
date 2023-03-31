package object

type Class struct {
	Name          string
	IsInterface   bool
	Methods       map[string]Method
	VariableCount int
}

func NewClass(name string, isInterface bool) *Class {
	var class Class
	class.Name = name
	class.IsInterface = isInterface
	class.Methods = make(map[string]Method)
	return &class
}
