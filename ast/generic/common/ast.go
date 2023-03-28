package common

import (
	"log"
	"reflect"
	"trident/utility"

	"gopkg.in/yaml.v3"
)

type Ast[T any] struct {
	astData     map[string]T
	packageName string
	classNames  []string
	funcNames   []string
}

func (ast *Ast[T]) init() {
	ast.astData = make(map[string]T)
	ast.classNames = make([]string, 0)
	ast.funcNames = make([]string, 0)
}

func (ast *Ast[T]) generateAstdataFromAstFile(path string) {
	bytes := utility.ReadLinesToBytes(path)

	err := yaml.Unmarshal([]byte(bytes), &ast.astData)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func (ast *Ast[T]) transverseAny(data any, filter func(string) bool, operate func(T) bool) string {
	va := reflect.ValueOf(data)
	var r string
	switch va.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		r = va.String()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		r = va.String()
	case reflect.Float32, reflect.Float64:
		r = va.String()
	case reflect.String:
		r = va.String()
	case reflect.Map:
		r = ast.transverseMap(data.(map[string]any), filter, operate)
	case reflect.Slice:
		r = ast.transverseSlice(data.([]T), filter, operate)

	default:
		panic("wrong basic types")
	}

	return r
}

func (ast *Ast[T]) transverseMap(a map[string]any, filter func(string) bool, operate func(T) bool) string {
	var r string
	for k, v := range a {
		// fmt.Println(k)
		if filter(k) {
			//operate bool deside when to stop transverse
			if operate(v.(T)) {
				return ""
			}
		}

		r = ast.transverseAny(v, filter, operate)
	}
	return r
}

func (ast *Ast[T]) transverseSlice(a []T, filter func(string) bool, operate func(T) bool) string {
	var r string
	for _, v := range a {
		r = ast.transverseAny(v, filter, operate)
	}
	return r
}

func (ast *Ast[T]) findPackageName(data any) string {
	var r string
	filter := func(s string) bool {
		if s == "Type=PackageDeclaration" {
			return true
		} else {
			return false
		}
	}

	operate := func(a T) bool {

		return true
	}

	ast.transverseAny(ast.astData, filter, operate)
	return r
}
