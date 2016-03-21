package models

type LoginForm struct {
    Username    string `validate:"required" name:"Username"`
    Password    string `validate:"required" name:"Password"`
}

