package controllers

import (
    "fmt"
)

// Constants of all the controllers and
// the URL their actions are routed to.
const (
    HOME = "/"
    
    ACCOUNT_DETAILS = "/account/details"
    ACCOUNT_CHANGE_PASSWORD = "/account/change_password"
)

// Functions of controllers which require
// a parameter (eg. a user profile requires
// a name passed to it)
var (
    EXAMPLE = func(id int) (string) {
        return fmt.Sprintf("/example/%d", id)
    }
)

// Templates for each controller action.
const (
    TMPL_HOME = "home"

    TMPL_ACCOUNT_DETAILS = "account/details"
    TMPL_ACCOUNT_CHANGE_PASSWORD = "account/change_password"
)

