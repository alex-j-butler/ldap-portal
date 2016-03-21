package models

type AccountSSHKeyForm struct {
    KeyName     string `validate:"required" name:"Key name"`
    Key         string `validate:"required" name:"Key"`
}

