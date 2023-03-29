package main

import (
	"fmt"
	"path/filepath"
	"trident/ast/generic/common"
)

func main() {
	// report.GenerateSortedCReport()
	fmt.Println("start")
	path := filepath.Join("D:\\", "ast", "1.ast")
	// path := filepath.Join("/Users", "ke", "go", "gotools", "trident", "ast", "generic", "testdata", "ast")
	ast := common.NewAstData(path)
	ast.GetPackageNames()
	ast.GetClassOrInterfaceNames()
	ast.GetMethodNames()
	ast.PrintData()
	fmt.Println("done")
}
