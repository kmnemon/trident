package object

type Interface struct {
	name    string
	methods map[string]Method
	lines   int
}

func NewInterface(name string) *Interface {
	var i Interface
	i.name = name
	i.methods = make(map[string]Method)
	return &i
}

func (i *Interface) GetName() string {
	return i.name
}

func (i *Interface) GetMethods() map[string]Method {
	return i.methods
}

func (i *Interface) GetLines() int {
	return i.lines
}
