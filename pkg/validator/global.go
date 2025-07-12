package validator

import "sync"

var globalValidator *CustomValidator
var once sync.Once

func Getvalidator() *CustomValidator {
	once.Do(func() {
		globalValidator = NewCustomValidator()
	})
	return globalValidator
}

// Untuk testing
func ResetValidator() {
    globalValidator = nil
    once = sync.Once{}
}