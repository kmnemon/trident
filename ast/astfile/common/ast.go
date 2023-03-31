package common

import (
	"errors"
	"log"
	"reflect"
	"trident/utility"

	"gopkg.in/yaml.v3"
)

type ast struct {
	astData map[string]any

	packageName string
}

func (a *ast) init() {
	a.astData = make(map[string]any)

}

func (a *ast) generateAstdataFromAstFile(path string) {
	bytes := utility.ReadLinesToBytes(path)

	err := yaml.Unmarshal([]byte(bytes), &a.astData)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func (a *ast) transverseAny(data any, filter func(string) bool, operate func(any) bool) string {
	if data == nil {
		return ""
	}

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
		r = a.transverseMap(data.(map[string]any), filter, operate)
	case reflect.Slice:
		r = a.transverseSlice(data.([]any), filter, operate)

	default:
		log.Println(data)
		log.Println(va)
		log.Fatal("wrong basic types")
	}

	return r
}

func (a *ast) transverseMap(data map[string]any, filter func(string) bool, operate func(any) bool) string {
	var r string
	for k, v := range data {
		// fmt.Println(k)
		if filter(k) {
			//operate bool decide when to stop transverse
			if operate(v) {
				return ""
			}
		}

		r = a.transverseAny(v, filter, operate)
	}
	return r
}

func (a *ast) transverseSlice(data []any, filter func(string) bool, operate func(any) bool) string {
	var r string
	for _, v := range data {
		r = a.transverseAny(v, filter, operate)
	}
	return r
}

func findNameOperate(data any, nameType string) (string, error) {
	if reflect.ValueOf(data).Kind() != reflect.Map {
		return "", errors.New("find name operate wrong, data is not map")
	}

	nameTypeFormat := "name(Type=" + nameType + ")"
	return data.(map[string]any)[nameTypeFormat].(map[string]any)["identifier"].(string), nil

}
