package models

type AccountChangePasswordForm struct {
    CurrentPassword     string `valid:"required"`
    NewPassword         string `valid:"required"`
    ConfirmPassword     string `valid:"required"`
}

