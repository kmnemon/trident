package object

type Method struct {
	name  string
	lines int
}

func NewMethod(name string) *Method {
	var m Method
	m.name = name
	return &m
}

func (m *Method) GetName() string {
	return m.name
}

func (m *Method) GetLines() int {
	return m.lines
}
