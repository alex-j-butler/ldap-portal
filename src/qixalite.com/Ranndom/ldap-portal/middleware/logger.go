package middleware

import (
    "time"

    "gopkg.in/macaron.v1"
    "qixalite.com/Ranndom/ldap-portal/modules/logging"
)

func Logger() macaron.Handler {
    return func(ctx *macaron.Context) {
        start := time.Now()

        rw := ctx.Resp.(macaron.ResponseWriter)
        ctx.Next()

        // 127.0.0.1 GET / - 200 in 2ms
        logging.HTTPLogger.Info("%s %s %s - %d in %v", ctx.RemoteAddr(), ctx.Req.Method, ctx.Req.RequestURI, rw.Status(), time.Since(start))
    }
}

