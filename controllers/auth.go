package controllers

import (
    "gopkg.in/macaron.v1"
)

const (
    AUTH_LOGIN = "/auth/login"
    AUTH_LOGOUT = "/auth/logout"
)

const (
    TMPL_AUTH_LOGIN = "auth/login"
)

func AuthLogin(ctx *macaron.Context) {
    ctx.Data["title"] = "Login"
    ctx.HTML(200, TMPL_AUTH_LOGIN)
}

func AuthLogout(ctx *macaron.Context) {

}

func POSTAuthLogin(ctx *macaron.Context) {
    ctx.Redirect(AUTH_LOGIN)
}

