package common

import "fmt"

func (ast *Ast) PrintData() {
	fmt.Println("~~~package~~~")
	fmt.Println(ast.packageName)
	fmt.Println("~~~interface~~~")
	fmt.Println(ast.interfaceNames)
	fmt.Println("~~~class~~~")
	fmt.Println(ast.classNames)
	fmt.Println("~~method~~~")
	fmt.Println(ast.methodNames)
}
