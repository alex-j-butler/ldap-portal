package controllers

import (
    "gopkg.in/macaron.v1"
)

func Home(ctx *macaron.Context) {
    ctx.Data["title"] = "Home"
    ctx.HTML(200, TMPL_HOME)
}

