package models

type AccountDetailsForm struct {
    GivenName   string `validate:"required" name:"Given name"`
    Surname     string `validate:"required" name:"Surname"`
}

