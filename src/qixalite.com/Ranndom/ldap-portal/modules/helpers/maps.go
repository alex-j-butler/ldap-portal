package helpers

import (
	"reflect"
)

func MapIndex(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}

	return -1
}

func MapInclude(vs []string, t string) bool {
	return MapIndex(vs, t) >= 0
}

func MapWhitelist(m *map[string]interface{}, whitelist []string) {
	for k, _ := range *m {
		if !MapInclude(whitelist, k) {
			delete(*m, k)
		}
	}
}

func UpdateStruct(s interface{}, m *map[string]interface{}) {
	reflectStruct := reflect.ValueOf(&s).Elem()

	for k, v := range *m {
		val := reflect.ValueOf(v)
		reflectStruct.FieldByName(k).Set(val)
	}
}

