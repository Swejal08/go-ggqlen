package utils

import (
	"reflect"

	goqu "github.com/doug-martin/goqu/v9"
)

func dereferenceStringPointer(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}

func ConvertInputFieldsToRecord(input interface{}, fieldMapper map[string]string) goqu.Record {
	// this function was created to map the graphql schema types into database columns

	record := goqu.Record{}

	inputValue := reflect.ValueOf(input)
	for i := 0; i < inputValue.NumField(); i++ {
		fieldName := inputValue.Type().Field(i).Name
		fieldValue := inputValue.Field(i)

		//Skip nil values and  id fields
		if fieldName != "ID" && !reflect.DeepEqual(fieldValue, reflect.Zero(inputValue.Field(i).Type())) {

			mappedName := fieldMapper[fieldName]
			if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() {

				record[mappedName] = fieldValue.Elem().Interface()
			}
		}
	}

	return record
}
