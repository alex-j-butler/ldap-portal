package controllers

import (
    "fmt"
)

// Constants of all the controllers and
// the URL their actions are routed to.
const (
    HOME = "/"
    
    ACCOUNT_DETAILS = "/account/details"
    ACCOUNT_SSH_KEYS = "/account/ssh_keys"
    ACCOUNT_CHANGE_PASSWORD = "/account/change_password"
)

// Functions of controllers which require
// a parameter (eg. a user profile requires
// a name passed to it)
var (
    ACCOUNT_EDIT_SSH_KEY = func(id int) (string) {
        return fmt.Sprintf("/account/ssh_keys/%d/edit", id)
    }
)

// Templates for each controller action.
const (
    TMPL_HOME = "home"

    TMPL_ACCOUNT_DETAILS = "account/details"
    TMPL_ACCOUNT_SSH_KEYS = "account/ssh_keys"
    TMPL_ACCOUNT_NEW_SSH_KEY = "account/new_ssh_key"
    TMPL_ACCOUNT_EDIT_SSH_KEY = "account/edit_ssh_key"
    TMPL_ACCOUNT_CHANGE_PASSWORD = "account/change_password"
)

