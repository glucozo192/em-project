package database

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func FieldMap(e any) ([]string, []any) {
	var fieldNames []string
	var fieldValues []any
	v := reflect.ValueOf(e).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fieldName := field.Tag.Get("db")
		fieldValue := v.Field(i).Addr().Interface()
		fieldNames = append(fieldNames, fieldName)
		fieldValues = append(fieldValues, fieldValue)
	}

	return fieldNames, fieldValues
}

func GetPlaceholdersForUnnest[T any](e T) string {
	var result []string

	v := reflect.ValueOf(e).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		typE := field.Type.String()
		switch typE {
		// for string value
		case
			reflect.TypeOf(pgtype.Text{}).String(),
			reflect.String.String():
			result = append(result, fmt.Sprintf("$%d::TEXT[]", i+1))

		// for float or double value
		case
			reflect.TypeOf(pgtype.Float4{}).String(),
			reflect.TypeOf(pgtype.Float8{}).String(),
			reflect.Float32.String(),
			reflect.Float64.String():
			result = append(result, fmt.Sprintf("$%d::FLOAT[]", i+1))

		// for timestamp value
		case
			reflect.TypeOf(pgtype.Timestamptz{}).String(),
			reflect.TypeOf(time.Time{}).String():
			result = append(result, fmt.Sprintf("$%d::TIMESTAMPTZ[]", i+1))
		}
	}

	return strings.Join(result, ", ")
}

func GetDataForUnnest[T any](data []T) []any {
	var (
		arr    [][]any
		result []any
	)
	if len(data) == 0 {
		return []any{}
	}
	numField := reflect.ValueOf(data[0]).Elem().NumField()
	for i := 0; i < numField; i++ {
		arr = append(arr, []any{})

	}

	for i := range data {
		v := reflect.ValueOf(data[i]).Elem()
		for j := 0; j < numField; j++ {
			value := v.Field(j).Addr().Interface()
			arr[j] = append(arr[j], value)
		}
	}

	for i := 0; i < numField; i++ {
		result = append(result, arr[i])
	}
	return result
}

func GeneratePlaceHolderForBulkUpsert(nuOfItems int, nuOfField int) string {
	var result []string
	for i := 1; i <= nuOfItems; i++ {
		var builder []string
		for j := 1; j <= nuOfField; j++ {
			builder = append(builder, fmt.Sprintf("$%d", (i-1)*nuOfField+j))
		}
		result = append(result, "("+strings.Join(builder, ", ")+")")
	}

	return strings.Join(result, ", ")
}
