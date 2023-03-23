package common

import (
	"log"
	"trident/utility"

	"gopkg.in/yaml.v3"
)

type Ast struct {
	astdata map[any]any
}

func (ast *Ast) init() {
	ast.astdata = make(map[any]any)
}

func (ast *Ast) generateAstdataFromAstFile(path string) {
	bytes := utility.ReadLinesToBytes(path)

	err := yaml.Unmarshal([]byte(bytes), &ast.astdata)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func (ast *Ast) GetAllFunctionsName() []string {
	return nil
}
