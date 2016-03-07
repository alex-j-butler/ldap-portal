package validation

import (
    validation "github.com/asaskevich/govalidator"
    "log"
)

func Validate(inter interface{}) (bool) {
    result, err := validation.ValidateStruct(inter)

    if err != nil {
        log.Printf("%s", err)
    }

    log.Printf("%s", result)

    return false
}

