package models

type LoginForm struct {
    Username    string `binding:"Required"`
    Password    string `binding:"Required"`
}

