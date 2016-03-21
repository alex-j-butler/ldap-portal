package validation

import (
    "reflect"
    "fmt"
    "gopkg.in/bluesuncorp/validator.v5"
)

var ValidationMessages = map[string]string{
    "required": "%s is required.",
    "alphanum": "%s must not contain symbols.",
}

func GetValidationMessage(fieldName string, fieldError string) string {
    return fmt.Sprintf(ValidationMessages[fieldError], fieldName)
}

func Validate(validateStruct interface{}) (bool, []string) {
    validate := validator.New("validate", validator.BakedInValidators)
    errs := validate.Struct(validateStruct)
    messages := make([]string, 1)

    if errs == nil {
        return true, nil
    }

    for _, err := range errs.Errors {
        var fieldName string
        t := reflect.TypeOf(validateStruct)
        field, _ := t.FieldByName(err.Field)

        tag := field.Tag.Get("name")

        if len(tag) == 0 {
            fieldName = err.Field
        } else {
            fieldName = tag
        }

        messages = append(messages, GetValidationMessage(fieldName, err.Tag))
    }

    if len(messages) == 0 {
        return true, nil
    } else {
        return false, messages
    }
}

