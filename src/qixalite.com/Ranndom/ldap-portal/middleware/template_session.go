package middleware

import (
    "gopkg.in/macaron.v1"
    "github.com/go-macaron/session"
)

const _VERSION = "1.0.0"

func Version() string {
    return _VERSION
}

// Macaron middleware to add the Session to the template data, allowing templates to
// call Session.Get() to retrieve session data.
func TemplateSessioner() macaron.Handler {
    return func(ctx *macaron.Context, sess session.Store) {
        ctx.Data["Session"] = sess
    }
}

