package common

import (
	"fmt"
	"log"
	"reflect"
	"trident/utility"

	"gopkg.in/yaml.v3"
)

type Ast[T any] struct {
	astdata map[string]T
}

func (ast *Ast[T]) init() {
	ast.astdata = make(map[string]T)
}

func (ast *Ast[T]) generateAstdataFromAstFile(path string) {
	bytes := utility.ReadLinesToBytes(path)

	err := yaml.Unmarshal([]byte(bytes), &ast.astdata)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func (ast *Ast[T]) transverseAny(a any, filter func(string) bool, operate func(T)) string {
	va := reflect.ValueOf(a)
	switch va.Kind() {
	case reflect.Map:
		ast.transverseMap(a.(map[string]any), filter, operate)
	case reflect.Slice:
		ast.transverseSlice(a.([]T), filter, operate)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return va.String()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return va.String()
	case reflect.Float32, reflect.Float64:
		return va.String()
	case reflect.String:
		return va.String()

	default:
		panic("less with wrong basic types")
	}

	return va.String()
}

func (ast *Ast[T]) transverseMap(a map[string]any, filter func(string) bool, operate func(T)) {
	for k, v := range a {
		fmt.Println(k)
		// if filter(k) {
		// 	operate(v)
		// }

		fmt.Println(ast.transverseAny(v, filter, operate))
	}
}

func (ast *Ast[T]) transverseSlice(a []T, filter func(string) bool, operate func(T)) {
	for _, v := range a {
		ast.transverseAny(v, filter, operate)
	}
}
