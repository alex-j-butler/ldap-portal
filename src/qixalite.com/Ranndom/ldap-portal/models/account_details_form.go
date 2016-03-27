package models

type AccountDetailsForm struct {
    Email       string `validate:"required" name:"Email"`
    GivenName   string `validate:"required" name:"Given name"`
    Surname     string `validate:"required" name:"Surname"`
}

