package arrays

import (
	"reflect"
	"strings"
)

func Contains(search interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == search {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(search)).IsValid() {
			return true
		}
	}
	return false
}

func ContainsIgnoreCase(search string, target []string) bool {
	if len(search) == 0 {
		return false
	}
	if len(target) == 0 {
		return false
	}
	search = strings.ToLower(search)
	for i := 0; i < len(target); i++ {
		if strings.ToLower(target[i]) == search {
			return true
		}
	}
	return false
}
