package models

type AccountSSHKeyForm struct {
    KeyName     string `binding:"Required"`
    Key         string `binding:"Required"`
}

