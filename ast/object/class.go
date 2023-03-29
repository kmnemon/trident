package object

type Class struct {
	name    string
	methods map[string]Method
	lines   int
}
