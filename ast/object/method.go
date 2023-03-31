package object

type Method struct {
	Name           string
	ParameterCount int
	Lines          int
}

func NewMethod(name string) *Method {
	var m Method
	m.Name = name
	return &m
}
