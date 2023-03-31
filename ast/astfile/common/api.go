package common

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"trident/ast/object"
)

func GenerateObjects(astfileDir string) *object.Project {
	p := object.NewProject(astfileDir)
	files, err := ioutil.ReadDir(astfileDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !strings.HasPrefix(file.Name(), ".") {
			NewAstFileAddToObject(filepath.Join(astfileDir, file.Name()), p)
		}

	}

	return p
}

func NewAstFileAddToObject(path string, p *object.Project) {
	var ast Ast
	ast.init()
	ast.generateAstdataFromAstFile(path)

	ast.findPackageName(p)
	ast.findClassOrInterfaceNames(p)
	ast.findMethodNames(p)
}
