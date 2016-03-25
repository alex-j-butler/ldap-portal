package middleware

import (
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/session"
    "github.com/go-macaron/csrf"
)

func CSRFToken() macaron.Handler {
    return func(ctx *macaron.Context, sess session.Store, x csrf.CSRF) {
        ctx.Data["CSRFToken"] = x.GetToken()
    }
}

