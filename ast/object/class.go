package object

type Class struct {
	name    string
	methods map[string]Method
	lines   int
}

func NewClass(name string) *Class {
	var class Class
	class.name = name
	class.methods = make(map[string]Method)
	return &class
}

func (c *Class) GetName() string {
	return c.name
}

func (c *Class) GetMethods() map[string]Method {
	return c.methods
}

func (c *Class) GetLines() int {
	return c.lines
}
