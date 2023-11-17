package database

import (
	"reflect"
)

func IsExistFieldInTable(dt any, target string) bool {
	t := reflect.TypeOf(dt)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagValue := field.Tag.Get("db")
		if tagValue == target {
			return true
		}
	}
	return false
}
