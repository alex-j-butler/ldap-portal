package controllers

import (
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/session"
)

const (
    ACCOUNT_DETAILS = "/account/details"
    ACCOUNT_SSH_KEYS = "/account/ssh_keys"
    ACCOUNT_CHANGE_PASSWORD = "/account/change_password"
)

const (
    TMPL_ACCOUNT_DETAILS = "account/details"
    TMPL_ACCOUNT_SSH_KEYS = "account/ssh_keys"
    TMPL_ACCOUNT_CHANGE_PASSWORD = "account/change_password"
)

func AccountDetails(ctx *macaron.Context) {
    ctx.Data["title"] = "Details"
    ctx.HTML(200, TMPL_ACCOUNT_DETAILS)
}

func AccountSSHKeys(ctx *macaron.Context) {
    ctx.Data["title"] = "SSH Keys"
    ctx.HTML(200, TMPL_ACCOUNT_SSH_KEYS)
}

func AccountChangePassword(ctx *macaron.Context) {
    ctx.Data["title"] = "Change password"
    ctx.HTML(200, TMPL_ACCOUNT_CHANGE_PASSWORD)
}

func POSTAccountDetails(ctx *macaron.Context, f *session.Flash) {
    f.Success("Example")

    ctx.Redirect(ACCOUNT_DETAILS)
}

func POSTAccountSSHKeys(ctx *macaron.Context) {
    ctx.Redirect(ACCOUNT_SSH_KEYS)
}

func POSTAccountChangePassword(ctx *macaron.Context) {
    ctx.Redirect(ACCOUNT_CHANGE_PASSWORD)
}

