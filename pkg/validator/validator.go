package validator

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validate *validator.Validate
}
type validUUID struct{
    UUID string `validate:"required,uuid"`
}
func NewCustomValidator() *CustomValidator {
	return &CustomValidator{
		validate: validator.New(),
	}
}

func (v *CustomValidator) Validate(i  interface{}) error {
	err := v.validate.Struct(i)
	if err != nil {
		return err
	}
	return nil
}

func (v *CustomValidator) isSQLInjection(input string) bool {
    patterns := []string{
        `(?i)select\s`, // SELECT keyword
        `(?i)insert\s`, // INSERT keyword
        `(?i)update\s`, // UPDATE keyword
        `(?i)delete\s`, // DELETE keyword
        `(?i)drop\s`,   // DROP keyword
        `(?i)union\s`,  // UNION keyword
        `--`,           // SQL comment
        `;`,            // Statement separator
        `'`,            // Single quote
        `"`,            // Double quote
    }

    for _, pattern := range patterns {
        matched, _ := regexp.MatchString(pattern, input)
        if matched {
            return true
        }
    }
    return false
}

// ValidateSQLInjection checks all string fields in a struct for SQL injection patterns
// ValidateJSONBody checks all string fields in a struct for SQL injection patterns
// and ensures no field contains only spaces.
func (v *CustomValidator) ValidateJSONBody(data interface{}) (string, string, bool) {
    val := reflect.ValueOf(data)

    // Ensure the input is a struct
    if val.Kind() == reflect.Ptr {
        val = val.Elem()
    }
    if val.Kind() != reflect.Struct {
        return "", "Invalid input type", false
    }

    // Iterate over all fields in the struct
    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        fieldType := val.Type().Field(i)

        // Only check string fields
        if field.Kind() == reflect.String {
            fieldValue := field.String()

            // Check for SQL Injection
            if v.isSQLInjection(fieldValue) {
                return fieldType.Name, "SQL injection detected", true
            }

            // Check for empty or only spaces
            if len(strings.TrimSpace(fieldValue)) == 0 {
                return fieldType.Name, "Field cannot be empty or only spaces", true
            }
        }
    }

    return "", "", false
}

// func (v *CustomValidator) GetQueryInteger(
//     re := regexp.MustCompile(`\d+`)
//     matches := re.FindStringSubmatch(query)
//     if len(matches) == 0 {
//         return 0, nil
//     }
//     var result int
//     _, err := fmt.Sscanf(matches[0], "%d", &result)
//     if err != nil {
//         return 0, err
//     }
//     return result, nil
// }

func(v *CustomValidator) IsValidUUID(uuid string) bool {
    // Regular expression to match UUID format
    req := validUUID{UUID: uuid}
    err := v.validate.Struct(req)
    if err != nil {
        return false
    }   
    return err == nil
}