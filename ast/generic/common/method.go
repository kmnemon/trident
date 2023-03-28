package common

import "strings"

func (ast *Ast[T]) GetAllMethodsName() []string {
	return nil
}

func (ast *Ast[T]) methodFilter(data string) bool {
	if strings.Contains(data, "Type=MethodDeclaration") {
		return true
	} else {
		return false
	}
}

func (ast *Ast[T]) methodOperate(a T) {

}
