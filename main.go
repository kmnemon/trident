package main

import (
	"fmt"
	"path/filepath"
	"trident/ast/astfile/common"
)

func main() {
	// report.GenerateSortedCReport()
	fmt.Println("start...")
	// path := filepath.Join(".", "ast")
	path := filepath.Join(".", "ast", "testdata")

	p := common.GenerateObjects(path)
	fmt.Println(p.ToString())

	fmt.Println("done")
}
