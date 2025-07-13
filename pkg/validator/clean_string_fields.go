package validator

import (
	"reflect"
	"strings"
)

// CleanStringFields trims spaces from all string fields in a struct
func CleanStringFields(data interface{}) {
    val := reflect.ValueOf(data)

    // Ensure the input is a pointer to a struct
    if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }
    if val.Kind() != reflect.Struct {
        return
    }

    // Iterate over all fields in the struct
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)

        // Only process string fields
        if field.Kind() == reflect.String {
            field.SetString(strings.TrimSpace(field.String()))
        }
    }
}