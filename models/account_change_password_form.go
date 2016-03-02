package models

type AccountChangePasswordForm struct {
    CurrentPassword     string `binding:"Required"`
    NewPassword         string `binding:"Required"`
    ConfirmPassword     string `binding:"Required"`
}

