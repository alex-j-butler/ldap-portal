package controllers

import (
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/session"
    "qixalite.com/Ranndom/ldap-portal/modules/helpers"
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

func AccountDetails(ctx *macaron.Context, f *session.Flash, sess session.Store) {
    if helpers.LoggedIn(ctx, sess) != true {
        f.Error("You must be logged in to access that!")
        ctx.Redirect(HOME)
        return
    }

    ctx.Data["title"] = "Details"
    ctx.HTML(200, TMPL_ACCOUNT_DETAILS)
}

func AccountSSHKeys(ctx *macaron.Context, f *session.Flash, sess session.Store) {
    if helpers.LoggedIn(ctx, sess) != true {
        f.Error("You must be logged in to access that!")
        ctx.Redirect(HOME)
        return
    }

    ctx.Data["title"] = "SSH Keys"
    ctx.HTML(200, TMPL_ACCOUNT_SSH_KEYS)
}

func AccountChangePassword(ctx *macaron.Context) {
    if helpers.LoggedIn(ctx, sess) != true {
        f.Error("You must be logged in to access that!")
        ctx.Redirect(HOME)
        return
    }

    ctx.Data["title"] = "Change password"
    ctx.HTML(200, TMPL_ACCOUNT_CHANGE_PASSWORD)
}

func POSTAccountDetails(ctx *macaron.Context, f *session.Flash, sess session.Store) {
    if helpers.LoggedIn(ctx, sess) != true {
        f.Error("You must be logged in to access that!")
        ctx.Redirect(HOME)
        return
    }

    ctx.Redirect(ACCOUNT_DETAILS)
}

func POSTAccountSSHKeys(ctx *macaron.Context, f *session.Flash, sess session.Store) {
    if helpers.LoggedIn(ctx, sess) != true {
        f.Error("You must be logged in to access that!")
        ctx.Redirect(HOME)
        return
    }

    ctx.Redirect(ACCOUNT_SSH_KEYS)
}

func POSTAccountChangePassword(ctx *macaron.Context, f *session.Flash, sess session.Store) {
    if helpers.LoggedIn(ctx, sess) != true {
        f.Error("You must be logged in to access that!")
        ctx.Redirect(HOME)
        return
    }

    ctx.Redirect(ACCOUNT_CHANGE_PASSWORD)
}

