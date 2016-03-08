package models

type AccountDetailsForm struct {
    GivenName   string `binding:"Required"`
    Surname     string `binding:"Required"`
}

