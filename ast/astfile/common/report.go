package common

import "fmt"

func (ast *Ast) PrintData() {
	fmt.Println("~~~package~~~")
	fmt.Println(ast.packageName)
	fmt.Println("~~~interface~~~")
	fmt.Println(ast.interfacesNames)
	fmt.Println("~~~class~~~")
	fmt.Println(ast.classesNames)
	fmt.Println("~~method~~~")
	fmt.Println(ast.methodsNames)
}
